package handler

import (
	"belajar-fiber/database"
	"belajar-fiber/model/entity"
	"belajar-fiber/model/request"
	"belajar-fiber/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {

	user := new(request.LoginRequest)
	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": err,
		})
	}
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.JSON(fiber.Map{
			"message": "failed request",
			"error":   errValidate.Error(),
		})
	}
	// validate email is found or not
	var users entity.User
	errs := database.DB.First(&users, "email = ?", user.Email).Error
	if errs != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "email not found",
		})
	}

	// compare password with hash and input

	isValid := utils.ComparePassword(user.Password, users.Password)
	if !isValid {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "password not valid",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "login success",
	})
}
