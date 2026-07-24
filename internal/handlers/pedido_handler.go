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

// PedidoHandler recibe las solicitudes HTTP relacionadas con pedidos.
// No trabaja directamente con la base de datos, sino mediante el service.
type PedidoHandler struct {
	servicio *services.PedidoService
}

// NuevoPedidoHandler crea el handler y recibe el PedidoService.
func NuevoPedidoHandler(s *services.PedidoService) *PedidoHandler {
	return &PedidoHandler{servicio: s}
}

// responderErrorPedido convierte los errores del negocio
// en códigos de respuesta HTTP.
func responderErrorPedido(w http.ResponseWriter, err error) {
	switch {

	// 422: la solicitud se entiende, pero sus datos no son válidos.
	case errors.Is(err, services.ErrDatosInvalidos),
		errors.Is(err, services.ErrReferenciaInvalida):

		RespondError(
			w,
			http.StatusUnprocessableEntity,
			err.Error(),
		)

	// 409: existe un conflicto con el estado actual.
	// Por ejemplo, stock insuficiente o pedido ya cancelado.
	case errors.Is(err, services.ErrStockInsuficiente),
		errors.Is(err, services.ErrEstadoInvalido):

		RespondError(
			w,
			http.StatusConflict,
			err.Error(),
		)

	// 404: el pedido solicitado no existe.
	case errors.Is(err, services.ErrNoEncontrado):

		RespondError(
			w,
			http.StatusNotFound,
			err.Error(),
		)

	// 500: error inesperado del servidor.
	default:
		RespondError(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
	}
}

// obtenerIDPedido obtiene el parámetro "id" de la URL
// y lo convierte de texto a número.
func obtenerIDPedido(r *http.Request) (uint, error) {

	// Ejemplo de URL: /api/v1/pedidos/5
	// chi.URLParam obtiene el valor "5".
	id, err := strconv.ParseUint(
		chi.URLParam(r, "id"),
		10,
		64,
	)

	// El ID no puede contener letras ni ser cero.
	if err != nil || id == 0 {
		return 0, errors.New("ID inválido")
	}

	return uint(id), nil
}

// Crear recibe los datos JSON de un nuevo pedido.
func (h *PedidoHandler) Crear(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Variable donde se guardarán los datos recibidos.
	var pedido models.Pedido

	// Convierte el JSON enviado en una estructura Pedido.
	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {

		// Si el JSON está mal escrito, responde 400.
		RespondError(
			w,
			http.StatusBadRequest,
			"JSON inválido",
		)
		return
	}

	// El service aplica las reglas del negocio:
	// referencias, stock, descuento y estado.
	if err := h.servicio.Crear(&pedido); err != nil {

		// Convierte el error del service en código HTTP.
		responderErrorPedido(w, err)
		return
	}

	// Si el pedido fue creado, responde 201.
	RespondJSON(
		w,
		http.StatusCreated,
		pedido,
	)
}

// Listar devuelve todos los pedidos registrados.
func (h *PedidoHandler) Listar(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Solicita la lista al service.
	lista, err := h.servicio.Listar()
	if err != nil {
		responderErrorPedido(w, err)
		return
	}

	// Si todo sale bien, responde 200 con la lista.
	RespondJSON(
		w,
		http.StatusOK,
		lista,
	)
}

// ObtenerPorID busca un pedido específico.
func (h *PedidoHandler) ObtenerPorID(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Obtiene y valida el ID de la URL.
	id, err := obtenerIDPedido(r)
	if err != nil {
		RespondError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	// Solicita el pedido al service.
	pedido, err := h.servicio.ObtenerPorID(id)
	if err != nil {

		// Si no existe, el error será convertido en 404.
		responderErrorPedido(w, err)
		return
	}

	// Si existe, responde 200 con los datos.
	RespondJSON(
		w,
		http.StatusOK,
		pedido,
	)
}

// Cancelar cambia el estado de un pedido pendiente a cancelado.
func (h *PedidoHandler) Cancelar(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Obtiene el ID de la URL.
	id, err := obtenerIDPedido(r)
	if err != nil {
		RespondError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	// El service verifica el estado del pedido
	// y devuelve la cantidad al stock.
	if err := h.servicio.Cancelar(id); err != nil {
		responderErrorPedido(w, err)
		return
	}

	// Si se canceló correctamente, responde 200.
	RespondJSON(
		w,
		http.StatusOK,
		map[string]string{
			"mensaje": "pedido cancelado",
		},
	)
}
