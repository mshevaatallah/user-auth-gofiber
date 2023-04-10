package handler

import (
	"belajar-fiber/database"
	"belajar-fiber/model/entity"
	"belajar-fiber/model/request"
	"belajar-fiber/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
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

	// generate token
	claims := jwt.MapClaims{}

	claims["email"] = users.Email
	claims["username"] = users.Username
	claims["phone"] = users.Phone
	claims["id"] = users.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	if user.Email == "mshevaatallah@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}
	token, errGenerate := utils.GenerateToken(&claims)
	if errGenerate != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "error generate token",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "login success",
		"token":   token,
	})
}
