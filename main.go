package main

import (
	// from iko.go
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hello User")
	})

	app.Get("/userauth")
	app.Get("/userauth")
	app.Get("/userauth")
	app.Get("/userauth")
	app.Get("/userauth")
	app.Get("/userauth")
	app.Get("/userauth")

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
