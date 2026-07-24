package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joancema/examen-panaderia/internal/handlers"
	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/services"
	"github.com/joancema/examen-panaderia/internal/storage"
)

// Este test verifica que stock insuficiente responda con 409.
func TestCrearPedidoConStockInsuficienteRespondeConflicto(
	t *testing.T,
) {
	// Repositorios en memoria.
	pedidos := storage.NuevoPedidoMemoria()
	productos := storage.NuevoProductoMemoria()
	clientes := storage.NuevoClienteMemoria()

	// El producto solo tiene una unidad disponible.
	producto := models.Producto{
		Nombre:         "Pan campesino",
		PrecioUnitario: 10,
		Stock:          1,
		Activo:         true,
	}

	// Cliente válido.
	cliente := models.Cliente{
		Nombre:   "Ana Pérez",
		Cedula:   "1310000001",
		Telefono: "0990000001",
	}

	// Guardamos el producto.
	if err := productos.Crear(&producto); err != nil {
		t.Fatalf(
			"no se pudo crear el producto: %v",
			err,
		)
	}

	// Guardamos el cliente.
	if err := clientes.Crear(&cliente); err != nil {
		t.Fatalf(
			"no se pudo crear el cliente: %v",
			err,
		)
	}

	// Creamos el service.
	servicio := services.NuevoPedidoService(
		pedidos,
		productos,
		clientes,
	)

	// Creamos el handler.
	handler := handlers.NuevoPedidoHandler(servicio)

	// Se solicitan 2 unidades,
	// pero el producto solamente tiene stock 1.
	body := fmt.Sprintf(
		`{"producto_id":%d,"cliente_id":%d,"cantidad":2}`,
		producto.ID,
		cliente.ID,
	)

	// Simula una solicitud HTTP POST.
	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/pedidos",
		strings.NewReader(body),
	)

	// Guarda la respuesta producida por el handler.
	recorder := httptest.NewRecorder()

	// Ejecuta directamente el método Crear.
	handler.Crear(recorder, request)

	// Debe responder 409 Conflict,
	// porque el producto existe, pero no tiene stock suficiente.
	if recorder.Code != http.StatusConflict {
		t.Errorf(
			"se esperaba status 409, se obtuvo %d. Body: %s",
			recorder.Code,
			recorder.Body.String(),
		)
	}
}
