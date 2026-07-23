package services

import (
	"strings"

	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/storage"
)

// ClienteService contiene las validaciones básicas de Cliente.
type ClienteService struct {
	repo storage.ClienteRepository
}

func NuevoClienteService(repo storage.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) Crear(c *models.Cliente) error {
	if c == nil ||
		strings.TrimSpace(c.Nombre) == "" ||
		strings.TrimSpace(c.Cedula) == "" ||
		strings.TrimSpace(c.Telefono) == "" {
		return ErrDatosInvalidos
	}

	return s.repo.Crear(c)
}

func (s *ClienteService) ObtenerPorID(id uint) (models.Cliente, error) {
	cliente, encontrado := s.repo.ObtenerPorID(id)
	if !encontrado {
		return models.Cliente{}, ErrNoEncontrado
	}

	return cliente, nil
}

func (s *ClienteService) Listar() ([]models.Cliente, error) {
	return s.repo.Listar()
}
