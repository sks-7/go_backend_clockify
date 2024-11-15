package routes

import (
	"clockfy_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterTagRoutes(app *fiber.App) {
	app.Get("/tag", controllers.GetAllTagsController)
	app.Get("tag/:id", controllers.FindTagByIdController)
	app.Post("tag/new", controllers.CreateTagController)
	app.Get("tagupdate/:id", controllers.TagupdateController)
	app.Get("tagdelete/:id", controllers.TagdeleteController)

}
