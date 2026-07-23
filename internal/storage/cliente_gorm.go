package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-panaderia/internal/models"
)

// ClienteGORM implementa ClienteRepository mediante GORM.
type ClienteGORM struct {
	db *gorm.DB
}

func NuevoClienteGORM(db *gorm.DB) *ClienteGORM {
	return &ClienteGORM{db: db}
}

func (r *ClienteGORM) Crear(c *models.Cliente) error {
	return r.db.Create(c).Error
}

func (r *ClienteGORM) ObtenerPorID(id uint) (models.Cliente, bool) {
	var cliente models.Cliente

	if err := r.db.First(&cliente, id).Error; err != nil {
		return models.Cliente{}, false
	}

	return cliente, true
}

func (r *ClienteGORM) Listar() ([]models.Cliente, error) {
	var lista []models.Cliente

	err := r.db.Find(&lista).Error
	return lista, err
}
