package nested_example

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/crud"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type nestedHandler struct {
	pai   *crud.GormHandler[StructPai]
	filho *crud.GormHandler[StructFilho]
}

func RegisterCustomRoutes(app *fiber.App, db *gorm.DB) {
	r := nestedHandler{
		crud.NewGormRepository[StructPai](db),
		crud.NewGormRepository[StructFilho](db),
	}
	routes := app.Group("/pai")
	routes.Get("/duplicar", r.DuplicarFilhosMasSemRelacaoComPai)
}

func (n nestedHandler) DuplicarFilhosMasSemRelacaoComPai(c *fiber.Ctx) error {

	pai, err := n.pai.Limit(1).GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao buscar registros",
			"error":   err.Error(),
		})
	}
	pai[0].Filhos = append(pai[0].Filhos, pai[0].Filhos...)
	filhosDuplicados, err := n.filho.Create(pai[0].Filhos)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao criar registros",
			"error":   err.Error(),
		})
	}

	return c.JSON(filhosDuplicados)
}
