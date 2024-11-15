package main

import (
	"clockfy_backend/dbconfig"
	"clockfy_backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	if err := dbconfig.InitDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbconfig.CloseDB()

	app := fiber.New()

	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Running Go Server!")
	})

	routes.RegisterClientRoutes(app)
	routes.RegisterTagRoutes(app)
	routes.RegisterProjectRoutes(app)

	log.Fatal(app.Listen(":9002"))
}
