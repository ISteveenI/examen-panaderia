package storage

import (
	"errors"

	"gorm.io/gorm"

	"github.com/joancema/examen-panaderia/internal/models"
)

// TAREA (CP1): Implemente ClienteGORM contra la interfaz ClienteRepository.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos:
//     los tests de acceptance/ compilan contra ellos.
//   - Guíese por ProductoGORM (producto_gorm.go): es el mismo patrón.
type ClienteGORM struct {
	db *gorm.DB
}

func NuevoClienteGORM(db *gorm.DB) *ClienteGORM {
	return &ClienteGORM{db: db}
}

func (r *ClienteGORM) Crear(c *models.Cliente) error {
	// TODO: implementar.
	return errors.New("TODO: implementar Crear")
}

func (r *ClienteGORM) ObtenerPorID(id uint) (models.Cliente, bool) {
	// TODO: implementar.
	return models.Cliente{}, false
}

func (r *ClienteGORM) Listar() ([]models.Cliente, error) {
	// TODO: implementar.
	return nil, errors.New("TODO: implementar Listar")
}
