package controller

import (
	"github.com/divyanshu050303/user_basic_operation/repository"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	Repo *repository.BookRepository
}

func (ctrl *BookController) CreateBook(c *fiber.Ctx) error {
	return nil
}
func (br *BookController) GetBooks(context *fiber.Ctx) error {
	return nil
}
func (br *BookController) GetBookById(context *fiber.Ctx) error {
	return nil
}
func (br *BookController) DeleteBook(context *fiber.Ctx) error {
	return nil
}
func (br *BookController) UpdateBook(context *fiber.Ctx) error {
	return nil
}
