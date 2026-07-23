// ARCHIVO BLOQUEADO — NO MODIFICAR
//
// Las 5 reglas de negocio se verifican aquí usando los repositorios EN MEMORIA
// (ya implementados en el repo base) como fakes. Así, estos tests miden solo
// la lógica de su PedidoService, sin depender de su implementación GORM.
package acceptance

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/services"
	"github.com/joancema/examen-panaderia/internal/storage"
)

type entornoReglas struct {
	svc          *services.PedidoService
	productos *storage.ProductoMemoria
	clientes     *storage.ClienteMemoria
	pedidos   *storage.PedidoMemoria
	principal      models.Producto
	ana          models.Cliente
}

func nuevoEntornoReglas(t *testing.T) entornoReglas {
	t.Helper()
	hm := storage.NuevoProductoMemoria()
	cm := storage.NuevoClienteMemoria()
	am := storage.NuevoPedidoMemoria()

	principal := models.Producto{Nombre: "Pan campesino", PrecioUnitario: 8.5, Stock: 10, Activo: true}
	require.NoError(t, hm.Crear(&principal))
	ana := models.Cliente{Nombre: "Ana Zambrano", Cedula: "1310000001", Telefono: "0990000001"}
	require.NoError(t, cm.Crear(&ana))

	return entornoReglas{
		svc:          services.NuevoPedidoService(am, hm, cm),
		productos: hm,
		clientes:     cm,
		pedidos:   am,
		principal:      principal,
		ana:          ana,
	}
}

// R1: no se crea un pedido si el producto no existe o está inactivo,
// o si el cliente no existe.
func TestCP2_R1_ReferenciasValidas(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Pedido{ProductoID: 99999, ClienteID: e.ana.ID, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un producto inexistente debe devolver ErrReferenciaInvalida")

	extra := models.Producto{Nombre: "Torta de bodas", PrecioUnitario: 15, Stock: 3, Activo: false}
	require.NoError(t, e.productos.Crear(&extra))
	a = models.Pedido{ProductoID: extra.ID, ClienteID: e.ana.ID, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un producto INACTIVO debe devolver ErrReferenciaInvalida")

	a = models.Pedido{ProductoID: e.principal.ID, ClienteID: 99999, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un cliente inexistente debe devolver ErrReferenciaInvalida")
}

// R2: la cantidad no puede superar el stock disponible.
func TestCP2_R2_StockInsuficiente(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Pedido{ProductoID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 11}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrStockInsuficiente,
		"pedir 11 unidades con stock 10 debe devolver ErrStockInsuficiente")
}

// R3: Total = Cantidad x PrecioUnitario, con 10% de descuento desde 5 unidades.
func TestCP2_R3_CalculoTotal(t *testing.T) {
	e := nuevoEntornoReglas(t)

	sinDescuento := models.Pedido{ProductoID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 3}
	require.NoError(t, e.svc.Crear(&sinDescuento),
		"crear un pedido válido no debe devolver error")
	require.InDelta(t, 25.50, sinDescuento.Total, 0.001,
		"3 x 8.50 = 25.50 (sin descuento)")
	require.Equal(t, models.EstadoPendiente, sinDescuento.Estado,
		"un pedido recién creado debe quedar en estado PENDIENTE")

	conDescuento := models.Pedido{ProductoID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 5}
	require.NoError(t, e.svc.Crear(&conDescuento))
	require.InDelta(t, 38.25, conDescuento.Total, 0.001,
		"5 x 8.50 = 42.50, con 10% de descuento = 38.25")
}

// R4: solo se puede cancelar un pedido en estado PENDIENTE.
func TestCP2_R4_CancelarSoloPendiente(t *testing.T) {
	e := nuevoEntornoReglas(t)

	entregado := models.Pedido{
		ProductoID: e.principal.ID,
		ClienteID:     e.ana.ID,
		Cantidad:      1,
		Estado:        models.EstadoEntregado,
		Total:         8.5,
	}
	require.NoError(t, e.pedidos.Crear(&entregado))
	require.ErrorIs(t, e.svc.Cancelar(entregado.ID), services.ErrEstadoInvalido,
		"cancelar un pedido ENTREGADO debe devolver ErrEstadoInvalido")

	require.ErrorIs(t, e.svc.Cancelar(99999), services.ErrNoEncontrado,
		"cancelar un pedido inexistente debe devolver ErrNoEncontrado")
}

// R5: al crear se descuenta el stock; al cancelar, se repone.
func TestCP2_R5_ReposicionStock(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Pedido{ProductoID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 3}
	require.NoError(t, e.svc.Crear(&a))

	h, ok := e.productos.ObtenerPorID(e.principal.ID)
	require.True(t, ok)
	require.Equal(t, uint(7), h.Stock,
		"al crear un pedido de 3 unidades, el stock debe bajar de 10 a 7")

	require.NoError(t, e.svc.Cancelar(a.ID), "cancelar un pedido PENDIENTE debe funcionar")

	cancelado, ok := e.pedidos.ObtenerPorID(a.ID)
	require.True(t, ok)
	require.Equal(t, models.EstadoCancelado, cancelado.Estado,
		"tras cancelar, el pedido debe quedar en estado CANCELADO")

	h, ok = e.productos.ObtenerPorID(e.principal.ID)
	require.True(t, ok)
	require.Equal(t, uint(10), h.Stock,
		"al cancelar, las 3 unidades deben reponerse al stock (7 -> 10)")
}
