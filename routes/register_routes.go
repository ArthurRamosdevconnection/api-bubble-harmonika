package routes

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/crud"
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/teste"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAllRoutes(app *fiber.App, db *gorm.DB) {

	crud.RegisterCrudRoute[teste.TesteModel](app, db)
	teste.RegisterRoutes(app, db)

}
