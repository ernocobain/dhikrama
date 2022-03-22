package user

import (
	"dhikrama.com/web/src/model/database"
	"dhikrama.com/web/src/model/entity/users"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerCreate(c *fiber.Ctx) error {
	userReq := new(users.UserRequest)

	if err := c.BodyParser(userReq); err != nil {
		return err
	}

	errors := ValidateUser(*userReq)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newUser := users.User{
		Name:  userReq.Name,
		Email: userReq.Email,
		Phone: userReq.Phone,
	}

	err := database.DB.Create(&newUser).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"massage": "succes",
		"data":    newUser,
	})
}
