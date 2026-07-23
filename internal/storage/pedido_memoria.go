// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-panaderia/internal/models"
)

// PedidoMemoria implementa PedidoRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type PedidoMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Pedido
	nextID uint
}

func NuevoPedidoMemoria() *PedidoMemoria {
	return &PedidoMemoria{datos: make(map[uint]models.Pedido), nextID: 1}
}

func (r *PedidoMemoria) Crear(a *models.Pedido) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	a.ID = r.nextID
	r.nextID++
	r.datos[a.ID] = *a
	return nil
}

func (r *PedidoMemoria) ObtenerPorID(id uint) (models.Pedido, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.datos[id]
	return a, ok
}

func (r *PedidoMemoria) Listar() ([]models.Pedido, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Pedido, 0, len(r.datos))
	for _, a := range r.datos {
		lista = append(lista, a)
	}
	return lista, nil
}

func (r *PedidoMemoria) Actualizar(a *models.Pedido) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[a.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[a.ID] = *a
	return nil
}
