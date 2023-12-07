package middleware

import (
	"gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(ctx *fiber.Ctx) error{
	token := ctx.Get("x-token")
	if token == ""{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "unauthorized",
		})
	}

	_, err := utils.VerifyToken(token)
	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "unauthorized",
		})
	}

	ctx.Locals("userInfo", claims)
	ctx.Locals("role", claims["name"])

	return ctx.Next();
}

func PermissionCreate(ctx *fiber.Ctx) error {

	return ctx.Next();
}