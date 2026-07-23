// ARCHIVO BLOQUEADO — NO MODIFICAR
package services

import (
	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/storage"
)

// ProductoService contiene la lógica de negocio de la Entidad A.
// Está completo: úselo como ejemplo de cómo un service valida datos,
// devuelve errores de dominio y delega la persistencia al repository.
type ProductoService struct {
	repo storage.ProductoRepository
}

func NuevoProductoService(repo storage.ProductoRepository) *ProductoService {
	return &ProductoService{repo: repo}
}

func (s *ProductoService) Crear(h *models.Producto) error {
	if h.Nombre == "" || h.PrecioUnitario <= 0 {
		return ErrDatosInvalidos
	}
	return s.repo.Crear(h)
}

func (s *ProductoService) ObtenerPorID(id uint) (models.Producto, error) {
	h, ok := s.repo.ObtenerPorID(id)
	if !ok {
		return models.Producto{}, ErrNoEncontrado
	}
	return h, nil
}

func (s *ProductoService) Listar() ([]models.Producto, error) {
	return s.repo.Listar()
}
