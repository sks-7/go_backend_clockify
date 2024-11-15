package routes

import (
	"clockfy_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterProjectRoutes(app *fiber.App) {
	app.Get("/project", controllers.GetAllProjectsController)
	app.Post("project/new", controllers.CreateProjectController)
	app.Get("/:id", controllers.FindProjectByIdController)
	app.Get("/projectdelete/:id", controllers.DeleteProjectController)
	app.Get("/projectupdate/:id", controllers.UpdateProjectController)
}
