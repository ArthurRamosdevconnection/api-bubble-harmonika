package teste

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/src/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (r handler) DeleteTeste(c *fiber.Ctx) error {
	id := c.Params("id")
	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Id inválido",
			"error":   err.Error(),
		})
	}
	err = r.db.Delete(&models.Teste{}, id_int).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "erro no deletar testes",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "teste deletado",
	})
}
