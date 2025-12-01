package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    app.Get("/api/hello", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "Hello from Fiber!",
        })
    })

    app.Listen(":8080")
}

