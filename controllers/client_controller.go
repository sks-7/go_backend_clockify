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

func GetAllClient(c *fiber.Ctx) error {

	ctx := context.Background()

	clientdata, err := dbconfig.Client.Client.FindMany().Exec(ctx)

	if err != nil {

		return c.Status(500).JSON(fiber.Map{
			"error": "error finding Clients",
		})

	}

	return c.Status(200).JSON(fiber.Map{
		"data": clientdata,
	})

}

func CreateClientController(c *fiber.Ctx) error {
	ctx := context.Background()

	var ClientReq model.Client
	if err := c.BodyParser(&ClientReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	data, err := dbconfig.Client.Client.CreateOne(
		db.Client.Name.Set(ClientReq.Name),
		db.Client.Address.Set(ClientReq.Address),
	).Exec(ctx)

	if err != nil {

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

func FindClientByIdController(c *fiber.Ctx) error {
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
	Client, err := dbconfig.Client.Client.FindFirst(db.Client.ID.Equals(id)).Exec(ctx)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error finding Client",
			"details": err.Error(),
		})
	}

	// Return the found Client
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Client found successfully",
		"data": Client,
	})
}

// // tag delete

func DeleteClientController(c *fiber.Ctx) error {
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
	Client, err := dbconfig.Client.Client.FindUnique(db.Client.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error finding Client",
			"details": err.Error(),
		})
	}

	// Return the found Client
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Tag delete successfully",
		"data": Client,
	})
}

// Client  upadter for

func UpdateClientController(c *fiber.Ctx) error {
	// Parse the ID from the URL parameter
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Create a context
	ctx := context.Background()

	var ClientReq model.Client

	// Bind the JSON request body to the tag struct
	err = c.BodyParser(&ClientReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Fetch the tag from the database
	Client, err := dbconfig.Client.Client.FindUnique(db.Client.ID.Equals(id)).Update(

		db.Client.Name.Set(ClientReq.Name),
		db.Client.Address.Set(ClientReq.Address),
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
		"data": Client,
	})

}
