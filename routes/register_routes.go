package routes

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/crud"
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/nested_example"
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/teste"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAllRoutes(app *fiber.App, db *gorm.DB) {

	crud.RegisterCrudRoute[teste.TesteModel](app, db, "/teste")
	crud.RegisterCrudRoute[nested_example.StructPai](app, db, "/pai")
	crud.RegisterCrudRoute[nested_example.StructFilho](app, db, "/filho")
	nested_example.RegisterCustomRoutes(app, db)
	teste.RegisterCustomRoutes(app, db)

}
