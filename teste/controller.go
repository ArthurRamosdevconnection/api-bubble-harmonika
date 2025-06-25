package teste

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/crud"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InterfaceTeste interface {
	crud.InterfaceCrud[TesteModel]
	//Outras funções estariam declaradas aqui
}

type Repo struct {
	InterfaceTeste
	db *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	tRepo := &Repo{
		db: db,
	}
	routes := app.Group("/api/v1/teste")
	routes.Get("/", tRepo.FindByValue)
}
