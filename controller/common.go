package controller

import "github.com/gofiber/fiber/v2"

func Homepage(c *fiber.Ctx) error {
	return c.SendString("welcome to the home page")
}
