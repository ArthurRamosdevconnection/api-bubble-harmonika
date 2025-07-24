package crud

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (r *GormHandler[T]) RouteCreate(c *fiber.Ctx) error {
	var toCreate []T
	err := c.BodyParser(&toCreate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao criar registro",
			"error":   err.Error(),
		})
	}
	created, err := r.Create(toCreate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao criar registro",
			"error":   err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Registro criado com sucesso",
		"data":    created,
	})
}
func (r *GormHandler[T]) RouteUpdate(c *fiber.Ctx) error {
	var toUpdate T
	err := c.BodyParser(&toUpdate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao criar registro",
			"error":   err.Error(),
		})
	}
	updated, err := r.Update(toUpdate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao criar registro",
			"error":   err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Registro criado com sucesso",
		"data":    updated,
	})
}

func (r *GormHandler[T]) RouteGetAll(c *fiber.Ctx) error {
	result, err := r.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao buscar registros",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(result)
}

func (r *GormHandler[T]) RouteFilters(c *fiber.Ctx) error {
	query := r
	for _, key := range GetJSONFieldNames[T]() {
		fmt.Println(key)
		value := c.Query(key)
		if value == "" {
			continue
		}
		query.Where(key, value)
	}
	result, err := query.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao buscar registros",
			"error":   err.Error(),
		})
	}
	id1, err := r.Where("id", 1).GetAll()
	fmt.Println(id1)
	return c.Status(200).JSON(result)

}

func (r *GormHandler[T]) RouteDelete(c *fiber.Ctx) error {
	idToDelete := c.Params("id")
	if idToDelete == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "id inválido",
		})
	}
	idInInt, err := strconv.Atoi(idToDelete)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "erro no id",
		})
	}

	deleted, err := r.Delete(idInInt)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Erro ao deletar registro",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Registro deletado com sucesso",
		"data":    deleted,
	})
}
