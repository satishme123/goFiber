package controller

import (
	"simple-api-with-go-fiber/database"
	models "simple-api-with-go-fiber/model"

	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func createResponseUser(user models.User) UserSerializer {
	return UserSerializer{user.ID, user.FirstName, user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := createResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)
	responseUsers := []UserSerializer{}
	for _, user := range users {
		responseUsers = append(responseUsers, createResponseUser(user))
	}
	return c.Status(200).JSON(responseUsers)
}
