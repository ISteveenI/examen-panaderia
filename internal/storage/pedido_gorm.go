package storage

import (
	"errors"

	"gorm.io/gorm"

	"github.com/joancema/examen-panaderia/internal/models"
)

// TAREA (CP2): Implemente PedidoGORM contra la interfaz PedidoRepository.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Guíese por ProductoGORM: es el mismo patrón con una entidad distinta.
//   - Recuerde: aquí NO va lógica de negocio. Solo persistencia.
type PedidoGORM struct {
	db *gorm.DB
}

func NuevoPedidoGORM(db *gorm.DB) *PedidoGORM {
	return &PedidoGORM{db: db}
}

func (r *PedidoGORM) Crear(a *models.Pedido) error {
	// TODO: implementar.
	return errors.New("TODO: implementar Crear")
}

func (r *PedidoGORM) ObtenerPorID(id uint) (models.Pedido, bool) {
	// TODO: implementar.
	return models.Pedido{}, false
}

func (r *PedidoGORM) Listar() ([]models.Pedido, error) {
	// TODO: implementar.
	return nil, errors.New("TODO: implementar Listar")
}

func (r *PedidoGORM) Actualizar(a *models.Pedido) error {
	// TODO: implementar.
	return errors.New("TODO: implementar Actualizar")
}
