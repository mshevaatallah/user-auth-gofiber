package handler

import (
	"belajar-fiber/database"
	"belajar-fiber/model/entity"
	"belajar-fiber/model/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return ctx.JSON(fiber.Map{
			"message": "error",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    users,
	})

}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserRequest)
	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "email exist",
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

	newUser := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.JSON(fiber.Map{
			"message": "error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    newUser,
	})

}
