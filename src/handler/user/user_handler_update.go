package user

import (
	"dhikrama.com/web/src/model/database"
	"dhikrama.com/web/src/model/entity/users"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userReq := new(users.UserRequest)

	if err := ctx.BodyParser(userReq); err != nil {
		return err
	}

	errorRes := ValidateUser(*userReq)
	if errorRes != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorRes)
	}

	var userC users.User

	userId := ctx.Params("id")

	err := database.DB.First(&userC, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userC.Name = userReq.Name
	userC.Email = userReq.Email
	userC.Phone = userReq.Phone
	errC := database.DB.Save(&userC).Error
	if errC != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errC.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    userC,
	})
}
