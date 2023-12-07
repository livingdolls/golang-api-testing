package route

import (
	"gofiber/handler"
	"gofiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Post("/login", handler.Login)

	r.Get("/users", middleware.UserMiddleware, handler.UserHanlerGetAll)
	r.Post("/users", handler.UserHandlerCreate)
	r.Get("users/:id", handler.UserHanlderGetById)
	r.Put("users/:id", handler.UserHandlerUpdate)
	r.Delete("users/:id", handler.UserHandlerDelete)

	// r.Post("/book",)
}