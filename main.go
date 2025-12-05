package main

import (
    "github.com/gofiber/fiber/v2"
    // "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()
    

    // API ROUTES

    app.Get("/api/hello", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "Hello API"})
    })

    //  Gets the static files from go 
    app.Static("/", "./frontend/dist")

    // tells the browser to use react routing 
    app.Get("*", func(c *fiber.Ctx) error {
        return c.SendFile("./frontend/dist/index.html")
    })

    app.Listen(":3000")
}
