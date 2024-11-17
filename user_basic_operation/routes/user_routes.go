package routes

import (
	"github.com/divyanshu050303/user_basic_operation/controller"
	"github.com/divyanshu050303/user_basic_operation/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := &repository.UserRepository{DB: db}

	userContoller := &controller.UserController{Repo: userRepository}
	api := app.Group("/api/user")

	api.Post("/createUser", userContoller.CreateUser)
	api.Get("/getUser", userContoller.GetUsers)
	api.Get("/getUserById/:id", userContoller.GetUserById)
	api.Delete("/deletUser/:id", userContoller.DeleteUser)
	api.Put("/updateUser", userContoller.UpdateUser)
}
