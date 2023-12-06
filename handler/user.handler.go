package handler

import (
	"gofiber/database"
	"gofiber/model/entity"
	"gofiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHanlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	err := database.DB.Find(&users).Error

	if err != nil {
		log.Println(err)
	}

	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest);

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()

	errValidate := validate.Struct(user)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status": false,
			"error" : errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Adress: user.Adress,
		Phone: user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error;

	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Message" : "failed to store data",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": true,
		"data" : newUser,
	})
}

func UserHanlderGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id");

	var user entity.User

	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status" : false,
			"message" : "User not found",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status" : true,
		"data" : user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest);

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status" : false,
			"message" : "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id");
	err := database.DB.First(&user, "id = ?", userId).Error;

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status" : false,
			"message" : "user not found",
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}

	user.Adress = userRequest.Adress
	user.Phone = userRequest.Phone

	errUpdate := database.DB.Save(&user).Error;

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status" : false,
			"message" : "internal server error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status" : true,
		"data" : user,
	})

}