package routers

import 
(
    "github.com/gofiber/fiber/v2"
    "alexshelto/url_shorten_service/Controllers"
)

func SetUpRouters() *fiber.App {

    app := fiber.New()

    app.Post("/api/v1/p", Controller.CreateHashedUrlV1)
    app.Get("/p/:id", Controller.GetHashedUrl)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    return app
}



