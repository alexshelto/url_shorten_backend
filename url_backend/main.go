package main

import "github.com/gofiber/fiber/v2"

func create_shortened_url(c *fiber.Ctx) error {
    return c.SendString("Create URL")
}


func main() {
    app := fiber.New()

    apiV1 := app.Group("/api/v1") // /api/v1
    apiV1.Post("/p", create_shortened_url)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}

