package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/services"
)

// ClienteHandler expone las operaciones HTTP de Cliente.
type ClienteHandler struct {
	servicio *services.ClienteService
}

func NuevoClienteHandler(s *services.ClienteService) *ClienteHandler {
	return &ClienteHandler{servicio: s}
}

func (h *ClienteHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente

	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := h.servicio.Crear(&cliente); err != nil {
		switch {
		case errors.Is(err, services.ErrDatosInvalidos):
			RespondError(
				w,
				http.StatusUnprocessableEntity,
				err.Error(),
			)
		default:
			RespondError(
				w,
				http.StatusInternalServerError,
				err.Error(),
			)
		}
		return
	}

	RespondJSON(w, http.StatusCreated, cliente)
}

func (h *ClienteHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	RespondJSON(w, http.StatusOK, lista)
}

func (h *ClienteHandler) ObtenerPorID(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := strconv.ParseUint(
		chi.URLParam(r, "id"),
		10,
		64,
	)

	if err != nil || id == 0 {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	cliente, err := h.servicio.ObtenerPorID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, services.ErrNoEncontrado):
			RespondError(
				w,
				http.StatusNotFound,
				err.Error(),
			)
		default:
			RespondError(
				w,
				http.StatusInternalServerError,
				err.Error(),
			)
		}
		return
	}

	RespondJSON(w, http.StatusOK, cliente)
}
