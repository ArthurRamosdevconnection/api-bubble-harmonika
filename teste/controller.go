package teste

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/crud"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type testeCustomRepo struct {
	*crud.GormHandler[TesteModel]
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	r := testeCustomRepo{
		crud.NewGormRepository[TesteModel](db),
	}
	routes := app.Group("/teste")
	routes.Get("/", r.HelloWorld)
}

func (t testeCustomRepo) HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
