package teste

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/models"
	"github.com/gofiber/fiber/v2"
)

func (r handler) UpdateTeste(c *fiber.Ctx) error {
	var testes []models.Teste
	err := r.db.Save(&testes).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "erro no encontrar testes",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello World!",
	})
}
