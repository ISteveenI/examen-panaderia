// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/services"
)

// ProductoHandler expone la Entidad A por HTTP.
// Está completo: observe cómo decodifica el body, llama al service y
// MAPEA los errores de dominio a status codes. Ese mapeo es exactamente
// lo que usted debe replicar en sus propios handlers.
type ProductoHandler struct {
	servicio *services.ProductoService
}

func NuevoProductoHandler(s *services.ProductoService) *ProductoHandler {
	return &ProductoHandler{servicio: s}
}

func (h *ProductoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *ProductoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var producto models.Producto
	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if err := h.servicio.Crear(&producto); err != nil {
		switch {
		case errors.Is(err, services.ErrDatosInvalidos):
			RespondError(w, http.StatusUnprocessableEntity, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusCreated, producto)
}
