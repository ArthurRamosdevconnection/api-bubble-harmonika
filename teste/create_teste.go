package teste

import "github.com/gofiber/fiber/v2"

func (r *Repo) Create(c *fiber.Ctx) error {
	err := r.Create(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao criar registro",
			"error":   err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Registro criado com sucesso",
	})
}
