package main

import (
	"fmt"
	"log"
	"simple-api-with-go-fiber/database"
	"simple-api-with-go-fiber/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello word")
	database.ConnectDb()
	app := fiber.New()
	routers.SetupRouters(app)
	log.Fatal(app.Listen(":3000"))
}
