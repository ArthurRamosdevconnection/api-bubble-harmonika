package teste

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/src/models"
	"github.com/gofiber/fiber/v2"
)

func (r handler) CreateTeste(c *fiber.Ctx) error {
	var payload models.Teste
	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "modelo inválido",
			"error":   err.Error(),
		})
	}
	err = r.db.Create(&payload).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao criar registro",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "objeto criado com sucesso!",
	})
}
