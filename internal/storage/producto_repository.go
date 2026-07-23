// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-panaderia/internal/models"

// ProductoRepository define el contrato de persistencia de la Entidad A.
type ProductoRepository interface {
	Crear(h *models.Producto) error
	ObtenerPorID(id uint) (models.Producto, bool)
	Listar() ([]models.Producto, error)
	Actualizar(h *models.Producto) error
}
