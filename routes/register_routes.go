package routes

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/teste"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {

	teste.RegisterRoutes(app, db)

}
