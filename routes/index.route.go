package routes

import (
	"belajar-fiber/handler"
	"belajar-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Post("/login", handler.LoginHandler)
	r.Get("/users", middleware.AuthMiddleware, handler.UserHandlerGetAll)
	r.Post("/users", handler.UserHandlerCreate)
	r.Get("/users/:id", handler.UserHandlerGetById)
	r.Put("/users/:id", handler.UserHandlerUpdate)
	r.Put("/users/:id/email-update", handler.UserHandlerUpdateEmail)
	r.Delete("/users/:id", handler.UserHandlerDelete)

}
