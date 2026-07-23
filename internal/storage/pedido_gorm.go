package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-panaderia/internal/models"
)

// PedidoGORM implementa PedidoRepository mediante GORM.
type PedidoGORM struct {
	db *gorm.DB
}

func NuevoPedidoGORM(db *gorm.DB) *PedidoGORM {
	return &PedidoGORM{db: db}
}

func (r *PedidoGORM) Crear(a *models.Pedido) error {
	return r.db.Create(a).Error
}

func (r *PedidoGORM) ObtenerPorID(id uint) (models.Pedido, bool) {
	var pedido models.Pedido

	err := r.db.
		Preload("Producto").
		Preload("Cliente").
		First(&pedido, id).Error

	if err != nil {
		return models.Pedido{}, false
	}

	return pedido, true
}

func (r *PedidoGORM) Listar() ([]models.Pedido, error) {
	var lista []models.Pedido

	err := r.db.
		Preload("Producto").
		Preload("Cliente").
		Find(&lista).Error

	return lista, err
}

func (r *PedidoGORM) Actualizar(a *models.Pedido) error {
	return r.db.Save(a).Error
}
