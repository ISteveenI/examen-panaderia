// ARCHIVO BLOQUEADO — NO MODIFICAR
package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/joancema/examen-panaderia/internal/handlers"
	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/services"
	"github.com/joancema/examen-panaderia/internal/storage"
)

func main() {
	db, err := gorm.Open(sqlite.Open("panaderia.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Producto{},
		&models.Cliente{},
		&models.Pedido{},
	); err != nil {
		log.Fatalf("error en la migración: %v", err)
	}

	sembrarProductos(db)

	// Repositories (GORM)
	productoRepo := storage.NuevoProductoGORM(db)
	clienteRepo := storage.NuevoClienteGORM(db)
	pedidoRepo := storage.NuevoPedidoGORM(db)

	// Services
	productoSvc := services.NuevoProductoService(productoRepo)
	clienteSvc := services.NuevoClienteService(clienteRepo)
	pedidoSvc := services.NuevoPedidoService(pedidoRepo, productoRepo, clienteRepo)

	// Handlers + Router
	router := handlers.NuevoRouter(
		handlers.NuevoProductoHandler(productoSvc),
		handlers.NuevoClienteHandler(clienteSvc),
		handlers.NuevoPedidoHandler(pedidoSvc),
	)

	log.Println("API de la panadería escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

// sembrarProductos carga el catálogo inicial solo si la tabla está vacía.
// Los clientes y pedidos se crean vía API.
func sembrarProductos(db *gorm.DB) {
	var total int64
	db.Model(&models.Producto{}).Count(&total)
	if total > 0 {
		return
	}
	iniciales := []models.Producto{
		{Nombre: "Pan campesino", PrecioUnitario: 8.50, Stock: 10, Activo: true},
		{Nombre: "Torta de chocolate", PrecioUnitario: 6.00, Stock: 4, Activo: true},
		{Nombre: "Empanada de queso", PrecioUnitario: 5.00, Stock: 2, Activo: true},
		{Nombre: "Torta de bodas", PrecioUnitario: 15.00, Stock: 3, Activo: false},
	}
	for i := range iniciales {
		db.Create(&iniciales[i])
	}
}
