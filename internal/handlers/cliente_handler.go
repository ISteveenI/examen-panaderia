package handlers

import (
	"net/http"

	"github.com/joancema/examen-panaderia/internal/services"
)

// TAREA (CP1): Implemente ClienteHandler.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos:
//     routes.go (bloqueado) los registra y los tests httptest los atacan.
//   - Guíese por ProductoHandler para decodificar JSON y mapear errores:
//     ErrDatosInvalidos -> 422, ErrNoEncontrado -> 404.
//   - Para leer el {id} de la ruta: chi.URLParam(r, "id") y strconv.
type ClienteHandler struct {
	servicio *services.ClienteService
}

func NuevoClienteHandler(s *services.ClienteService) *ClienteHandler {
	return &ClienteHandler{servicio: s}
}

func (h *ClienteHandler) Crear(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 201 con el cliente creado.
	RespondError(w, http.StatusNotImplemented, "TODO: implementar")
}

func (h *ClienteHandler) Listar(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 200 con la lista.
	RespondError(w, http.StatusNotImplemented, "TODO: implementar")
}

func (h *ClienteHandler) ObtenerPorID(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar. Éxito -> 200; no existe -> 404.
	RespondError(w, http.StatusNotImplemented, "TODO: implementar")
}
