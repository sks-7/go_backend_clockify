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

func GetAllTagsController(c *fiber.Ctx) error {
	ctx := context.Background()
	dbPosts, err := dbconfig.Client.Tag.FindMany().Exec(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching posts",
		})
	}

	return c.JSON(dbPosts)
}

func CreateTagController(c *fiber.Ctx) error {
	ctx := context.Background()

	var tagtReq model.Tag
	if err := c.BodyParser(&tagtReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	data, err := dbconfig.Client.Tag.CreateOne(
		db.Tag.Name.Set(tagtReq.Name),
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

func FindTagByIdController(c *fiber.Ctx) error {
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
	tag, err := dbconfig.Client.Tag.FindFirst(db.Tag.ID.Equals(id)).Exec(ctx)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error finding tag",
			"details": err.Error(),
		})
	}

	// Return the found tag
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Tag found successfully",
		"data": tag,
	})
}

// tag delete

func TagdeleteController(c *fiber.Ctx) error {
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
	tag, err := dbconfig.Client.Tag.FindUnique(db.Tag.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error finding tag",
			"details": err.Error(),
		})
	}

	// Return the found tag
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Tag delete successfully",
		"data": tag,
	})
}

// tag  upadter for

func TagupdateController(c *fiber.Ctx) error {
	// Parse the ID from the URL parameter
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Create a context
	ctx := context.Background()

	var tagReq model.Tag

	// Bind the JSON request body to the tag struct
	err = c.BodyParser(&tagReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Fetch the tag from the database
	tag, err := dbconfig.Client.Tag.FindUnique(db.Tag.ID.Equals(id)).Update(

		db.Tag.Name.Set(tagReq.Name),
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
		"data": tag,
	})

}
