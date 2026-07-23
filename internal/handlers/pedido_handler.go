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

type PedidoHandler struct {
	servicio *services.PedidoService
}

func NuevoPedidoHandler(s *services.PedidoService) *PedidoHandler {
	return &PedidoHandler{servicio: s}
}

func responderErrorPedido(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, services.ErrDatosInvalidos),
		errors.Is(err, services.ErrReferenciaInvalida):
		RespondError(w, http.StatusUnprocessableEntity, err.Error())

	case errors.Is(err, services.ErrStockInsuficiente),
		errors.Is(err, services.ErrEstadoInvalido):
		RespondError(w, http.StatusConflict, err.Error())

	case errors.Is(err, services.ErrNoEncontrado):
		RespondError(w, http.StatusNotFound, err.Error())

	default:
		RespondError(w, http.StatusInternalServerError, err.Error())
	}
}

func obtenerIDPedido(r *http.Request) (uint, error) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id == 0 {
		return 0, errors.New("ID inválido")
	}

	return uint(id), nil
}

func (h *PedidoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var pedido models.Pedido

	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := h.servicio.Crear(&pedido); err != nil {
		responderErrorPedido(w, err)
		return
	}

	RespondJSON(w, http.StatusCreated, pedido)
}

func (h *PedidoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		responderErrorPedido(w, err)
		return
	}

	RespondJSON(w, http.StatusOK, lista)
}

func (h *PedidoHandler) ObtenerPorID(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerIDPedido(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	pedido, err := h.servicio.ObtenerPorID(id)
	if err != nil {
		responderErrorPedido(w, err)
		return
	}

	RespondJSON(w, http.StatusOK, pedido)
}

func (h *PedidoHandler) Cancelar(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerIDPedido(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.servicio.Cancelar(id); err != nil {
		responderErrorPedido(w, err)
		return
	}

	RespondJSON(
		w,
		http.StatusOK,
		map[string]string{"mensaje": "pedido cancelado"},
	)
}
