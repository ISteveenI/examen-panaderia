package handlers

import (
	"net/http"

	"github.com/joancema/examen-panaderia/internal/services"
)

// TAREA (CP3): Implemente PedidoHandler.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Mapeo de errores de dominio a status codes (los tests lo verifican):
//       ErrDatosInvalidos     -> 422 Unprocessable Entity
//       ErrReferenciaInvalida -> 422 Unprocessable Entity
//       ErrStockInsuficiente  -> 409 Conflict
//       ErrEstadoInvalido     -> 409 Conflict
//       ErrNoEncontrado       -> 404 Not Found
//       cualquier otro error  -> 500 Internal Server Error
type PedidoHandler struct {
	servicio *services.PedidoService
}

func NuevoPedidoHandler(s *services.PedidoService) *PedidoHandler {
	return &PedidoHandler{servicio: s}
}

func (h *PedidoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 201 con el pedido creado (incluido el total).
	RespondError(w, http.StatusNotImplemented, "TODO: implementar")
}

func (h *PedidoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 200 con la lista.
	RespondError(w, http.StatusNotImplemented, "TODO: implementar")
}

func (h *PedidoHandler) ObtenerPorID(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 200; no existe -> 404.
	RespondError(w, http.StatusNotImplemented, "TODO: implementar")
}

func (h *PedidoHandler) Cancelar(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 200; estado inválido -> 409; no existe -> 404.
	RespondError(w, http.StatusNotImplemented, "TODO: implementar")
}
