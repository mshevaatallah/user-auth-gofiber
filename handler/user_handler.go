package handler

import (
	"belajar-fiber/database"
	"belajar-fiber/model/entity"
	"belajar-fiber/model/request"
	"belajar-fiber/model/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
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
		return ctx.Status(404).JSON(fiber.Map{
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
		return ctx.Status(404).JSON(fiber.Map{
			"message": "eerror create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    newUser,
	})

}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user entity.User
	err := database.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "error",
		})
	}
	response := response.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    response,
	})

}
