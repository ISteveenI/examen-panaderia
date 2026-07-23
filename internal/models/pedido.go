package models

import "gorm.io/gorm"

// TAREA (CP1): Complete los campos de Pedido según lo que muestran las pantallas.
//
// Pistas de trabajo:
//   - Un Pedido referencia a un Producto y a un Cliente (claves foráneas).
//   - Recuerde el campo de estado (use las constantes de estados.go) y el total.
//   - Los tests de acceptance/ compilan contra los nombres EXACTOS de los campos.
type Pedido struct {
	gorm.Model
	// TODO: agregue aquí los campos.
}
