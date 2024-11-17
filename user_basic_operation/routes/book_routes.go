package routes

import (
	"github.com/divyanshu050303/user_basic_operation/controller"
	"github.com/divyanshu050303/user_basic_operation/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUpBookRoutes(app *fiber.App, db *gorm.DB) {
	bookRepo := &repository.BookRepository{DB: db}

	bookController := &controller.BookController{Repo: bookRepo}

	api := app.Group("/api/book")
	api.Post("/createBook", bookController.CreateBook)
	api.Get("/getBook", bookController.GetBooks)
	api.Get("/getBookById/:id", bookController.GetBookById)
	api.Delete("/deletBook/:id", bookController.DeleteBook)
	api.Put("/updateBook", bookController.UpdateBook)
}
