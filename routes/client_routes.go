package routes

import (
	"clockfy_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterClientRoutes(app *fiber.App) {
	app.Get("/client", controllers.GetAllClient)
	app.Get("client/:id", controllers.FindClientByIdController)
	app.Post("client/new", controllers.CreateClientController)
	app.Get("clientupdate/:id", controllers.UpdateClientController)
	app.Get("clientdelete/:id", controllers.DeleteClientController)
}
