package services

import (
	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/storage"
)

// TAREA (CP1): Implemente ClienteService.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Cliente no tiene reglas de negocio complejas: valide lo evidente según
//     las pantallas (campos obligatorios -> ErrDatosInvalidos) y delegue al
//     repository. Guíese por ProductoService.
type ClienteService struct {
	repo storage.ClienteRepository
}

func NuevoClienteService(repo storage.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) Crear(c *models.Cliente) error {
	// TODO: implementar.
	return ErrNoImplementado
}

func (s *ClienteService) ObtenerPorID(id uint) (models.Cliente, error) {
	// TODO: implementar.
	return models.Cliente{}, ErrNoImplementado
}

func (s *ClienteService) Listar() ([]models.Cliente, error) {
	// TODO: implementar.
	return nil, ErrNoImplementado
}
