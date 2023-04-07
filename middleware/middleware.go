package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(ctx *fiber.Ctx) error {

	token := ctx.Get("x-token")
	if token != "secret" {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	return ctx.Next()
}
