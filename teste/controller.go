package teste

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	r := &handler{db}

	routes := app.Group("/api/v1/teste")
	routes.Get("/", r.GetTeste)
	routes.Get("/", r.GetTesteById)
	routes.Post("/", r.CreateTeste)
}
