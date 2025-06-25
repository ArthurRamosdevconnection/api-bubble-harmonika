package teste

import "github.com/gofiber/fiber/v2"

func (r *Repo) FindByValue(c *fiber.Ctx) error {
	value := c.Query("value")
	r.db.Where("value like %?%", value)
	result, err := r.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao buscar o registro",
		})
	}
	return c.Status(200).JSON(result)
}
