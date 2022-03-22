package user

import (
	"log"

	"dhikrama.com/web/src/model/database"
	"dhikrama.com/web/src/model/entity/users"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerRead(c *fiber.Ctx) error {
	var user []users.User
	result := database.DB.Debug().Find(&user)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(user)
}
