package main

import (
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Template engine setup
	templates := template.Must(template.ParseGlob("views/*.html"))

	// Home Page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("/home/nelsonel/Projects/golang/Today/views/home.html", fiber.Map{})
	})

	// shop Page
	app.Get("/shop", func(c *fiber.Ctx) error {
		return c.Render("/home/nelsonel/Projects/golang/Today/views/shop.html", fiber.Map{})
	})

    // product Page
	app.Get("/shop", func(c *fiber.Ctx) error {
		return c.Render("/home/nelsonel/Projects/golang/Today/views/product.html", fiber.Map{})
	})

	// Contact Page
	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("/home/nelsonel/Projects/golang/Today/views/contact.html", fiber.Map{})
	})
    
	// Static files
	app.Static("/", "./public")

	// Middleware to render templates
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("templates", templates)
		return c.Next()
	})

	// Start the server
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
