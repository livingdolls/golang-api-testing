package route

import (
	"gofiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/users", handler.UserHanlerGetAll)
	r.Post("/users", handler.UserHandlerCreate)
	r.Get("users/:id", handler.UserHanlderGetById)
	r.Put("users/:id", handler.UserHandlerUpdate)
}