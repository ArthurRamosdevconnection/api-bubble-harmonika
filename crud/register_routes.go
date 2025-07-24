package crud

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterCrudRoute[T Model](app *fiber.App, db *gorm.DB) {
	r := NewGormRepository[T](db)

	routes := app.Group("/teste")
	routes.Get("/", r.RouteFilters)
	routes.Post("/", r.RouteCreate)
	routes.Put("/", r.RouteUpdate)

}
