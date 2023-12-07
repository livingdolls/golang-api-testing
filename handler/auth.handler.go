package handler

import (
	"gofiber/database"
	"gofiber/model/entity"
	"gofiber/model/request"
	"gofiber/utils"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New();
	errValidate := validate.Struct(loginRequest);

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "Failed",
			"error" : errValidate.Error(),
		})
	}

	var user entity.User;

	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message" : "email not found",
		})
	}

	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password);

	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "Wrong Password",
		})
	}

	// JWT
	claims := jwt.MapClaims{};
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["adress"] = user.Adress
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		log.Println(errToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "Wrong Password",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"token" : token,
	})
}