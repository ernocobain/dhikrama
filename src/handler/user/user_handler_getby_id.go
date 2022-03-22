package user

import (
	"dhikrama.com/web/src/model/database"
	"dhikrama.com/web/src/model/entity/users"
	"github.com/gofiber/fiber/v2"
)

func UserGetById(c *fiber.Ctx) error {
	userId := c.Params("id")

	var userC users.User

	err := database.DB.First(&userC, "id = ?", userId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{

			"message": err.Error(),
		})
	}

	// userResponse := user.UserResponse{
	// 	ID:        userC.ID,
	// 	Name:      userC.Name,
	// 	Email:     userC.Email,
	// 	Phone:     userC.Phone,
	// 	CreatedAt: userC.CreatedAt,
	// 	UpdatedAt: userC.UpdatedAt,
	// }

	return c.JSON(fiber.Map{
		"message": "succes",
		"data":    userC,
	})
}
