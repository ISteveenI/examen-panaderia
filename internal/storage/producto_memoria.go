// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-panaderia/internal/models"
)

// ProductoMemoria implementa ProductoRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type ProductoMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Producto
	nextID uint
}

func NuevoProductoMemoria() *ProductoMemoria {
	return &ProductoMemoria{datos: make(map[uint]models.Producto), nextID: 1}
}

func (r *ProductoMemoria) Crear(h *models.Producto) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	h.ID = r.nextID
	r.nextID++
	r.datos[h.ID] = *h
	return nil
}

func (r *ProductoMemoria) ObtenerPorID(id uint) (models.Producto, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	h, ok := r.datos[id]
	return h, ok
}

func (r *ProductoMemoria) Listar() ([]models.Producto, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Producto, 0, len(r.datos))
	for _, h := range r.datos {
		lista = append(lista, h)
	}
	return lista, nil
}

func (r *ProductoMemoria) Actualizar(h *models.Producto) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[h.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[h.ID] = *h
	return nil
}
