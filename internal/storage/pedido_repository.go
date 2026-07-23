// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-panaderia/internal/models"

// PedidoRepository define el contrato de persistencia de Pedido.
// Su implementación GORM (en pedido_gorm.go) debe satisfacer EXACTAMENTE
// estas firmas. Observe que el repositorio NO contiene lógica de negocio:
// las reglas (validaciones, cálculo del total, anulación) viven en el service.
type PedidoRepository interface {
	Crear(a *models.Pedido) error
	ObtenerPorID(id uint) (models.Pedido, bool)
	Listar() ([]models.Pedido, error)
	Actualizar(a *models.Pedido) error
}
