package services

import (
	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/storage"
)

// TAREA (CP2): Implemente PedidoService con las 5 reglas de negocio.
//
// Las reglas están A LA VISTA en las pantallas (carpeta pantallas/) y los
// tests de acceptance/reglas_test.go las verifican una por una. Devuelva los
// errores de dominio de errores.go: los tests los comprueban con errors.Is.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Observe que el service recibe TRES repositories: necesita consultar
//     Producto y Cliente para validar, y actualizar Producto al cancelar.
type PedidoService struct {
	pedidos   storage.PedidoRepository
	productos storage.ProductoRepository
	clientes     storage.ClienteRepository
}

func NuevoPedidoService(
	pedidos storage.PedidoRepository,
	productos storage.ProductoRepository,
	clientes storage.ClienteRepository,
) *PedidoService {
	return &PedidoService{
		pedidos:   pedidos,
		productos: productos,
		clientes:     clientes,
	}
}

// Crear registra un nuevo pedido aplicando R1, R2 y R3.
// TODO (R1): el producto debe existir y estar activo; el cliente debe existir.
// TODO (R2): la cantidad no puede superar el stock disponible del producto.
// TODO (R3): calcule el total (observe en las pantallas cuándo aplica descuento).
// TODO: al crear, el stock del producto se descuenta (mire la pantalla 01
// antes y después de crear un pedido; R5 es la operación inversa).
func (s *PedidoService) Crear(a *models.Pedido) error {
	// TODO: implementar.
	return ErrNoImplementado
}

func (s *PedidoService) ObtenerPorID(id uint) (models.Pedido, error) {
	// TODO: implementar.
	return models.Pedido{}, ErrNoImplementado
}

func (s *PedidoService) Listar() ([]models.Pedido, error) {
	// TODO: implementar.
	return nil, ErrNoImplementado
}

// Cancelar cancela un pedido aplicando R4 y R5.
// TODO (R4): solo se puede cancelar un pedido en estado PENDIENTE.
// TODO (R5): al cancelar, la cantidad se repone al stock del producto.
func (s *PedidoService) Cancelar(id uint) error {
	// TODO: implementar.
	return ErrNoImplementado
}
