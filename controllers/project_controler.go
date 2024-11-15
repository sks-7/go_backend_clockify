package controllers

import (
	"clockfy_backend/dbconfig"
	"clockfy_backend/model"
	"clockfy_backend/prisma/db"
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllProjectsController(c *fiber.Ctx) error {

	// var email = c.Query("email")

	ctx := context.Background()
	dbProjects, err := dbconfig.Client.Project.FindMany(
		// db.Project.Useremail.Equals(email),
	).Exec(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching Projects",
		})
	}

	return c.JSON(dbProjects)
}

func CreateProjectController(c *fiber.Ctx) error {
	ctx := context.Background()

	var projectReq model.Project

	log.Printf("Creating project",projectReq)
	if err := c.BodyParser(&projectReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}


	data, err := dbconfig.Client.Project.CreateOne(
		db.Project.Name.Set(projectReq.Name),
		db.Project.Tag.Set(projectReq.Tag),
		db.Project.Billable.Set(projectReq.Billable),
		db.Project.Useremail.Set(projectReq.Useremail),
	).Exec(ctx)

	if err != nil {

		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating user",
			"err":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  true,
		"data": data,
	})
}

func FindProjectByIdController(c *fiber.Ctx) error {
	// Parse the ID from the URL parameter
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Create a context
	ctx := context.Background()

	// Fetch the tag from the database
	project, err := dbconfig.Client.Project.FindFirst(db.Project.ID.Equals(id)).Exec(ctx)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error finding project",
			"details": err.Error(),
		})
	}

	// Return the found project
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "project found successfully",
		"data": project,
	})
}

// tag delete

func DeleteProjectController(c *fiber.Ctx) error {
	// Parse the ID from the URL parameter
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Create a context
	ctx := context.Background()

	// Fetch the tag from the database
	project, err := dbconfig.Client.Project.FindUnique(db.Project.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error finding project",
			"details": err.Error(),
		})
	}

	// Return the found project
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Tag delete successfully",
		"data": project,
	})
}

// tag  upadter for

func UpdateProjectController(c *fiber.Ctx) error {
	// Parse the ID from the URL parameter
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Create a context
	ctx := context.Background()

	var projectReq model.Project

	// Bind the JSON request body to the tag struct
	err = c.BodyParser(&projectReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Fetch the tag from the database
	project, err := dbconfig.Client.Project.FindUnique(db.Project.ID.Equals(id)).Update(

		db.Project.Name.Set(projectReq.Name),
		db.Project.Tag.Set(projectReq.Tag),
		db.Project.Billable.Set(projectReq.Billable),
		db.Project.Useremail.Set(projectReq.Useremail),
	).Exec(ctx)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error finding tag",
			"details": err.Error(),
		})
	}

	// Return the updated tag
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Tag updated successfully",
		"data": project,
	})

}
