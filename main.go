package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	users := app.Group("users")

	// User Endpoints
	users.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	users.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Money Transfer (Transaction) endpoint
	app.Post("/transfer-money", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
