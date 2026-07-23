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

func TestCrearPedidoConStockInsuficienteRespondeConflicto(t *testing.T) {
	pedidos := storage.NuevoPedidoMemoria()
	productos := storage.NuevoProductoMemoria()
	clientes := storage.NuevoClienteMemoria()

	producto := models.Producto{
		Nombre:         "Pan campesino",
		PrecioUnitario: 10,
		Stock:          1,
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

	handler := handlers.NuevoPedidoHandler(servicio)

	body := fmt.Sprintf(
		`{"producto_id":%d,"cliente_id":%d,"cantidad":2}`,
		producto.ID,
		cliente.ID,
	)

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/pedidos",
		strings.NewReader(body),
	)

	recorder := httptest.NewRecorder()

	handler.Crear(recorder, request)

	if recorder.Code != http.StatusConflict {
		t.Errorf(
			"se esperaba status 409, se obtuvo %d. Body: %s",
			recorder.Code,
			recorder.Body.String(),
		)
	}
}
