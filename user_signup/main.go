package main

import (
	"divyanshu050303/user_signup/models"
	"divyanshu050303/user_signup/routes"
	"divyanshu050303/user_signup/storage"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Cound not load the database")
	}
	err = models.MigrateUser(db)
	if err != nil {
		log.Fatal("could not migrate the datebase")
	}
	app := fiber.New()
	routes.SetUpUserRoutes(app, db)
	app.Listen(":8080")

}
