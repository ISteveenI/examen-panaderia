package services_test

import (
	"testing"

	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/services"
	"github.com/joancema/examen-panaderia/internal/storage"
)

func TestCrearPedidoAplicaDescuentoYDescuentaStock(t *testing.T) {
	pedidos := storage.NuevoPedidoMemoria()
	productos := storage.NuevoProductoMemoria()
	clientes := storage.NuevoClienteMemoria()

	producto := models.Producto{
		Nombre:         "Pan integral",
		PrecioUnitario: 10,
		Stock:          10,
		Activo:         true,
	}

	cliente := models.Cliente{
		Nombre:   "Ana Pérez",
		Cedula:   "1310000001",
		Telefono: "0990000001",
	}

	if err := productos.Crear(&producto); err != nil {
		t.Fatalf("no se pudo crear el producto: %v", err)
	}

	if err := clientes.Crear(&cliente); err != nil {
		t.Fatalf("no se pudo crear el cliente: %v", err)
	}

	servicio := services.NuevoPedidoService(
		pedidos,
		productos,
		clientes,
	)

	pedido := models.Pedido{
		ProductoID: producto.ID,
		ClienteID:  cliente.ID,
		Cantidad:   5,
	}

	if err := servicio.Crear(&pedido); err != nil {
		t.Fatalf("no se pudo crear el pedido: %v", err)
	}

	if pedido.Total != 45 {
		t.Errorf("se esperaba total 45, se obtuvo %.2f", pedido.Total)
	}

	if pedido.Estado != models.EstadoPendiente {
		t.Errorf("se esperaba estado PENDIENTE, se obtuvo %s", pedido.Estado)
	}

	actualizado, encontrado := productos.ObtenerPorID(producto.ID)
	if !encontrado {
		t.Fatal("producto no encontrado")
	}

	if actualizado.Stock != 5 {
		t.Errorf("se esperaba stock 5, se obtuvo %d", actualizado.Stock)
	}
}
