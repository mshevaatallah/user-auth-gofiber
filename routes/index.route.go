package routes

import (
	"belajar-fiber/handler"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Get("/users", handler.UserHandlerGetAll)
	r.Post("/users", handler.UserHandlerCreate)
	r.Get("/users/:id", handler.UserHandlerGetById)
	r.Put("/users/:id", handler.UserHandlerUpdate)
	r.Put("/users/:id/email-update", handler.UserHandlerUpdateEmail)

}
