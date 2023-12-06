package main

import (
	"gofiber/database"
	"gofiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit();
	app := fiber.New();

	route.RouteInit(app);

	app.Listen(":8181")
}