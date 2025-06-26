package teste

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/src/models"
	"github.com/gofiber/fiber/v2"
)

func (r handler) GetTeste(c *fiber.Ctx) error {
	var testes []models.Teste
	err := r.db.Find(&testes).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "erro no encontrar testes",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"testes": testes,
	})
}
