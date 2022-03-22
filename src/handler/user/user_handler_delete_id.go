package user

import (
	"dhikrama.com/web/src/model/database"
	"dhikrama.com/web/src/model/entity/users"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerDeleteId(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var userD users.User

	err := database.DB.Debug().First(&userD, "id =?", userId).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	errorD := database.DB.Debug().Delete(&userD).Error

	if errorD != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "User was deleted",
	})
}
