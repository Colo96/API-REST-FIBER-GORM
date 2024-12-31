package controllers

import (
	"api-rest-fiber-gorm/src/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var repo = &models.Repository{}

func CreateUser(context *fiber.Ctx) error {
	users := models.Users{}
	err := context.BodyParser(&users)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"},
		)
		return err
	}
	err = repo.DB.Create(&users).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create user"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "user has been created"},
	)
	return nil
}

func DeleteUser(context *fiber.Ctx) error {
	userModel := models.Users{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"},
		)
		return nil
	}
	err := repo.DB.Delete(userModel, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete book"},
		)
		return err.Error
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book has been deleted"},
	)
	return nil
}

func GetUserById(context *fiber.Ctx) error {
	userModel := &models.Users{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"},
		)
		return nil
	}
	err := repo.DB.Where("id = ?", id).First(userModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the user"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "user id fetched successfully", "data": userModel},
	)
	return nil
}

func GetUsers(context *fiber.Ctx) error {
	userModels := &[]models.Users{}
	err := repo.DB.Find(userModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the users"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "users fetched successfully", "data": userModels},
	)
	return nil
}

func UpdateUser(context *fiber.Ctx) error {
	users := models.Users{}
	err := context.BodyParser(&users)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Invalid request body"},
		)
		return err
	}

	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "ID cannot be empty"},
		)
		return nil
	}

	var existingUser models.Users
	err = repo.DB.Where("id = ?", id).First(&existingUser).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "User not found"},
		)
		return err
	}

	err = repo.DB.Model(&existingUser).Updates(users).Error
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Could not update user"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "User updated successfully", "data": existingUser},
	)
	return nil
}
