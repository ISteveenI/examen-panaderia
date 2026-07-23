package models

import "gorm.io/gorm"

// Pedido registra la compra de un producto realizada por un cliente.
type Pedido struct {
	gorm.Model

	ProductoID uint     `gorm:"not null;index" json:"producto_id"`
	Producto   Producto `gorm:"foreignKey:ProductoID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"producto"`

	ClienteID uint    `gorm:"not null;index" json:"cliente_id"`
	Cliente   Cliente `gorm:"foreignKey:ClienteID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"cliente"`

	Cantidad uint    `gorm:"not null" json:"cantidad"`
	Estado   string  `gorm:"size:20;not null;default:PENDIENTE" json:"estado"`
	Total    float64 `gorm:"not null" json:"total"`
}
