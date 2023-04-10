package middleware

import (
	"belajar-fiber/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {

	token := ctx.Get("x-token")

	if token == "" {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	_, err := utils.VerifyToken(token)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	//if token != "secret" {
	//	return ctx.Status(401).JSON(fiber.Map{
	//		"message": "unauthorized",
	//	})
	//}

	return ctx.Next()
}
