package main

import (
	"log"
	"os"

	"github.com/divyanshu050303/user_basic_operation/models"
	"github.com/divyanshu050303/user_basic_operation/routes"
	"github.com/divyanshu050303/user_basic_operation/storages"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storages.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := storages.NewConnection(config)
	if err != nil {
		log.Fatal("Could Not load datebase")
	}
	err = models.Migrate(db)
	if err != nil {
		log.Fatal("Count not migrate the databse")
	}
	app := fiber.New()

	routes.SetUpBookRoutes(app, db)
	routes.SetUserRoutes(app, db)
	app.Listen(":8080")
}
