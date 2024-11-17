package controller

import (
	"net/http"

	"github.com/divyanshu050303/user_basic_operation/models"
	"github.com/divyanshu050303/user_basic_operation/repository"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	Repo *repository.BookRepository
}

func (ctrl *BookController) CreateBook(c *fiber.Ctx) error {
	book := models.BookModels{
		ID: uuid.New().String(),
	}
	err := c.BodyParser(&book)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = ctrl.Repo.DB.Create(&book).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not add the book"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book has been added successfully"})
	return nil
}
func (br *BookController) GetBooks(context *fiber.Ctx) error {
	book := &[]models.BookModels{}
	err := br.Repo.DB.Find(book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get books"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book fetched successfully", "data": book})
	return nil
}
func (br *BookController) GetBookById(context *fiber.Ctx) error {
	id := context.Params("id")
	book := &models.BookModels{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id connot be empty"})
		return nil
	}
	err := br.Repo.DB.Where("id=?", id).Find(book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Book fetched successfully", "data": book})
	return nil
}
func (br *BookController) DeleteBook(context *fiber.Ctx) error {
	id := context.Params("id")
	book := &models.BookModels{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id connot be empty"})
		return nil
	}
	err := br.Repo.DB.Delete(book, id).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not Delete the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Book Delete successfully"})
	return nil
}
func (br *BookController) UpdateBook(context *fiber.Ctx) error {
	// Get the book ID from the URL parameters
	id := context.Params("id")

	// Find the book in the database
	var book models.BookModels
	err := br.Repo.DB.First(&book, id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Book not found")
	}

	// Parse the request body into a BookModels struct
	var updatedBook models.BookModels
	err = context.BodyParser(&updatedBook)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Update the book in the database
	book.Author = updatedBook.Author
	book.Title = updatedBook.Title
	book.Publication = updatedBook.Publication
	err = br.Repo.DB.Save(&book).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update book")
	}

	// Return a success response
	return context.JSON(fiber.Map{"message": "Book updated successfully"})
}
