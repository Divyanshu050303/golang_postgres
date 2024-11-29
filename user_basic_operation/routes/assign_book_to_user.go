package routes

import (
	"github.com/divyanshu050303/user_basic_operation/controller"
	"github.com/divyanshu050303/user_basic_operation/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupAssignBookToUserRoutes(app *fiber.App, db *gorm.DB) {
	assignUserToBookRepo := &repository.AssignBookToUser{DB: db}
	assignBookToUserCotroller := &controller.AssignBookToUserController{Repo: assignUserToBookRepo}
	api := app.Group("/api/assignBookToUser")
	api.Post("/assignBookToUser", assignBookToUserCotroller.AssignBookToUser)
	api.Get("/getBookByUserId/:userId", assignBookToUserCotroller.GetBookByUserId)
	api.Get("/getBookByBookId/:bookId", assignBookToUserCotroller.GetBookByBookId)
	api.Post("/returnBook/:bookId", assignBookToUserCotroller.ReturnBook)
}
