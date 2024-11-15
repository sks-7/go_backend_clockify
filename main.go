package main

import (
    "clockfy_backend/dbconfig"
    "clockfy_backend/routes"
    "log"
    "os" 

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/joho/godotenv"
)

func main() {

    godotenv.Load() // Load the environment variables from the .env file

    // Connect to the database
    if err := dbconfig.InitDB(); err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer dbconfig.CloseDB()

    app := fiber.New()

    // Middleware
    app.Use(cors.New())

    // Root route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Running Go Server!")
    })

    // Register routes
    routes.RegisterClientRoutes(app)
    routes.RegisterTagRoutes(app)
    routes.RegisterProjectRoutes(app)

    // Get the port from environment variables or use 9002 by default
    port := os.Getenv("PORT")
    if port == "" {
        port = "9002" 
    }

    // Start the server
    log.Fatal(app.Listen(":" + port))
}
