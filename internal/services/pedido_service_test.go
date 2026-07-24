package services_test

import (
	"testing"

	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/services"
	"github.com/joancema/examen-panaderia/internal/storage"
)

// Este test comprueba el descuento y la reducción del stock.
func TestCrearPedidoAplicaDescuentoYDescuentaStock(t *testing.T) {

	// Se utilizan repositorios en memoria como fakes.
	// Así no necesitamos una base de datos real.
	pedidos := storage.NuevoPedidoMemoria()
	productos := storage.NuevoProductoMemoria()
	clientes := storage.NuevoClienteMemoria()

	// Producto con precio 10 y stock 10.
	producto := models.Producto{
		Nombre:         "Pan integral",
		PrecioUnitario: 10,
		Stock:          10,
		Activo:         true,
	}

	// Cliente válido.
	cliente := models.Cliente{
		Nombre:   "Ana Pérez",
		Cedula:   "1310000001",
		Telefono: "0990000001",
	}

	// Guardamos el producto en memoria.
	if err := productos.Crear(&producto); err != nil {
		t.Fatalf(
			"no se pudo crear el producto: %v",
			err,
		)
	}

	// Guardamos el cliente en memoria.
	if err := clientes.Crear(&cliente); err != nil {
		t.Fatalf(
			"no se pudo crear el cliente: %v",
			err,
		)
	}

	// Creamos el service usando los repositorios fake.
	servicio := services.NuevoPedidoService(
		pedidos,
		productos,
		clientes,
	)

	// Se solicitan 5 unidades.
	// Desde 5 unidades se aplica 10 % de descuento.
	pedido := models.Pedido{
		ProductoID: producto.ID,
		ClienteID:  cliente.ID,
		Cantidad:   5,
	}

	// Ejecutamos el método Crear del service.
	if err := servicio.Crear(&pedido); err != nil {
		t.Fatalf(
			"no se pudo crear el pedido: %v",
			err,
		)
	}

	// 5 unidades por 10 = 50.
	// Con 10 % de descuento, el total debe ser 45.
	if pedido.Total != 45 {
		t.Errorf(
			"se esperaba total 45, se obtuvo %.2f",
			pedido.Total,
		)
	}

	// Todo pedido nuevo debe quedar PENDIENTE.
	if pedido.Estado != models.EstadoPendiente {
		t.Errorf(
			"se esperaba estado PENDIENTE, se obtuvo %s",
			pedido.Estado,
		)
	}

	// Buscamos nuevamente el producto para revisar su stock.
	actualizado, encontrado := productos.ObtenerPorID(
		producto.ID,
	)

	if !encontrado {
		t.Fatal("producto no encontrado")
	}

	// Tenía 10 unidades y se compraron 5.
	// El stock final debe ser 5.
	if actualizado.Stock != 5 {
		t.Errorf(
			"se esperaba stock 5, se obtuvo %d",
			actualizado.Stock,
		)
	}
}
