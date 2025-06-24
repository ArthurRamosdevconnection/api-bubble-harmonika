package teste

import "github.com/gofiber/fiber/v2"

func (r handler) GetTeste(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello World!",
	})
}
