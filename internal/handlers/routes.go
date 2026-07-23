// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NuevoRouter registra todas las rutas de la API. Este archivo es el
// contrato HTTP del examen: los tests httptest de acceptance/ atacan
// exactamente estas rutas.
func NuevoRouter(
	productos *ProductoHandler,
	clientes *ClienteHandler,
	pedidos *PedidoHandler,
) http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/productos", func(r chi.Router) {
			r.Get("/", productos.Listar)
			r.Post("/", productos.Crear)
		})

		r.Route("/clientes", func(r chi.Router) {
			r.Get("/", clientes.Listar)
			r.Post("/", clientes.Crear)
			r.Get("/{id}", clientes.ObtenerPorID)
		})

		r.Route("/pedidos", func(r chi.Router) {
			r.Get("/", pedidos.Listar)
			r.Post("/", pedidos.Crear)
			r.Get("/{id}", pedidos.ObtenerPorID)
			r.Post("/{id}/cancelar", pedidos.Cancelar)
		})
	})

	return r
}
