package handler

import "github.com/gofiber/fiber/v2"

func CarCreate(c *fiber.Ctx) error {
	type requestBody struct {
		Number      string `json:"number"`
		Description string `json:"description"`
		Phone       string `json:"phone"`
		IsDefault   bool   `json:"is_default"`
	}
	return nil
}

func CarUpdate(c *fiber.Ctx) error {
	return nil
}

func CarDelete(c *fiber.Ctx) error {
	return nil
}
