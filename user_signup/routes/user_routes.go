package routes

import (
	"divyanshu050303/user_signup/controller"
	"divyanshu050303/user_signup/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUpUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := &repository.UserRepository{DB: db}
	userController := &controller.UserController{Repo: userRepository}

	api := app.Group("/api/user")
	api.Post("/signup", userController.SignupUser)
	api.Post("/login", userController.LoginUser)
	api.Get("/getUser", userController.GetUsers)
	api.Post("/forgotPassword", userController.ForgotPassword)
	api.Post("/resetPassword", userController.ForgotPassword)
}
