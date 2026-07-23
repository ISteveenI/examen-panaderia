package services

import (
	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/storage"
)

type PedidoService struct {
	pedidos   storage.PedidoRepository
	productos storage.ProductoRepository
	clientes  storage.ClienteRepository
}

func NuevoPedidoService(
	pedidos storage.PedidoRepository,
	productos storage.ProductoRepository,
	clientes storage.ClienteRepository,
) *PedidoService {
	return &PedidoService{
		pedidos:   pedidos,
		productos: productos,
		clientes:  clientes,
	}
}

// Crear aplica las reglas R1, R2, R3 y parte de R5.
func (s *PedidoService) Crear(a *models.Pedido) error {
	if a == nil || a.ProductoID == 0 ||
		a.ClienteID == 0 || a.Cantidad == 0 {
		return ErrDatosInvalidos
	}

	// R1: el producto debe existir y estar activo.
	producto, existe := s.productos.ObtenerPorID(a.ProductoID)
	if !existe || !producto.Activo {
		return ErrReferenciaInvalida
	}

	// R1: el cliente debe existir.
	_, existe = s.clientes.ObtenerPorID(a.ClienteID)
	if !existe {
		return ErrReferenciaInvalida
	}

	// R2: la cantidad no puede superar el stock.
	if a.Cantidad > producto.Stock {
		return ErrStockInsuficiente
	}

	// R3: cantidad por precio unitario.
	a.Total = float64(a.Cantidad) * producto.PrecioUnitario

	// R4:descuento del 10 % desde 5 unidades.
	if a.Cantidad >= 5 {
		a.Total = a.Total * 0.90
	}

	// Todo pedido nuevo comienza pendiente.
	a.Estado = models.EstadoPendiente

	// R5: descontar del stock las unidades solicitadas.
	producto.Stock -= a.Cantidad

	if err := s.productos.Actualizar(&producto); err != nil {
		return err
	}

	return s.pedidos.Crear(a)
}

func (s *PedidoService) ObtenerPorID(id uint) (models.Pedido, error) {
	pedido, existe := s.pedidos.ObtenerPorID(id)
	if !existe {
		return models.Pedido{}, ErrNoEncontrado
	}

	return pedido, nil
}

func (s *PedidoService) Listar() ([]models.Pedido, error) {
	return s.pedidos.Listar()
}

// Cancelar aplica R4 y R5.
func (s *PedidoService) Cancelar(id uint) error {
	pedido, existe := s.pedidos.ObtenerPorID(id)
	if !existe {
		return ErrNoEncontrado
	}

	// R4: solamente se cancela si está pendiente.
	if pedido.Estado != models.EstadoPendiente {
		return ErrEstadoInvalido
	}

	producto, existe := s.productos.ObtenerPorID(pedido.ProductoID)
	if !existe {
		return ErrReferenciaInvalida
	}

	// R5: devolver al stock las unidades del pedido.
	producto.Stock += pedido.Cantidad

	if err := s.productos.Actualizar(&producto); err != nil {
		return err
	}

	pedido.Estado = models.EstadoCancelado

	return s.pedidos.Actualizar(&pedido)
}
