// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-panaderia/internal/models"
)

// ProductoGORM implementa ProductoRepository sobre GORM.
// Esta implementación está completa: úsela como plantilla para ClienteGORM
// y PedidoGORM, que usted debe implementar.
type ProductoGORM struct {
	db *gorm.DB
}

func NuevoProductoGORM(db *gorm.DB) *ProductoGORM {
	return &ProductoGORM{db: db}
}

func (r *ProductoGORM) Crear(h *models.Producto) error {
	return r.db.Create(h).Error
}

func (r *ProductoGORM) ObtenerPorID(id uint) (models.Producto, bool) {
	var h models.Producto
	if err := r.db.First(&h, id).Error; err != nil {
		return models.Producto{}, false
	}
	return h, true
}

func (r *ProductoGORM) Listar() ([]models.Producto, error) {
	var lista []models.Producto
	err := r.db.Find(&lista).Error
	return lista, err
}

func (r *ProductoGORM) Actualizar(h *models.Producto) error {
	return r.db.Save(h).Error
}
