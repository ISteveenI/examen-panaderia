// ARCHIVO BLOQUEADO — NO MODIFICAR
//
// Carpeta acceptance/: estos tests SON la rúbrica del examen.
// Ejecútelos cuantas veces quiera con:
//
//	go test ./acceptance/... -v
//
// Para ver solo un checkpoint:
//
//	go test ./acceptance/... -v -run TestCP1
//	go test ./acceptance/... -v -run TestCP2
//	go test ./acceptance/... -v -run TestCP3
//
// NOTA: hasta que usted complete los campos de los modelos (CP1), este
// paquete no compila. Los errores de compilación le indican exactamente
// qué campos faltan.
package acceptance

import (
	"net/http"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joancema/examen-panaderia/internal/handlers"
	"github.com/joancema/examen-panaderia/internal/models"
	"github.com/joancema/examen-panaderia/internal/services"
	"github.com/joancema/examen-panaderia/internal/storage"
)

// nuevaDB abre una base SQLite en memoria y migra los tres modelos.
func nuevaDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("no se pudo abrir SQLite en memoria: %v", err)
	}
	if err := db.AutoMigrate(
		&models.Producto{},
		&models.Cliente{},
		&models.Pedido{},
	); err != nil {
		t.Fatalf("la migración de los modelos falló: %v", err)
	}
	return db
}

// nuevoRouterCompleto arma el cableado completo de la API sobre la base dada:
// repositories GORM -> services -> handlers -> router.
func nuevoRouterCompleto(t *testing.T, db *gorm.DB) http.Handler {
	t.Helper()
	productoRepo := storage.NuevoProductoGORM(db)
	clienteRepo := storage.NuevoClienteGORM(db)
	pedidoRepo := storage.NuevoPedidoGORM(db)
	return handlers.NuevoRouter(
		handlers.NuevoProductoHandler(services.NuevoProductoService(productoRepo)),
		handlers.NuevoClienteHandler(services.NuevoClienteService(clienteRepo)),
		handlers.NuevoPedidoHandler(services.NuevoPedidoService(pedidoRepo, productoRepo, clienteRepo)),
	)
}
