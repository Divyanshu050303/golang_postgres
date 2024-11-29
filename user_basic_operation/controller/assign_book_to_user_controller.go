package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/divyanshu050303/user_basic_operation/models"
	"github.com/divyanshu050303/user_basic_operation/repository"
	"github.com/gofiber/fiber/v2"
)

type AssignBookToUserController struct {
	Repo *repository.AssignBookToUser
}

func (ctrl *AssignBookToUserController) AssignBookToUser(c *fiber.Ctx) error {
	var assignBook models.AssignBookToUserModel
	err := c.BodyParser(&assignBook)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "invalid request body"})
	}
	var user models.UserModel
	err = ctrl.Repo.DB.First(&user, "id=?", assignBook.UserID).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "user not found"})
	}
	var book models.BookModels
	err = ctrl.Repo.DB.First(&book, "id=?", assignBook.BookID).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "book not found"})
	}
	newAssign := models.AssignBookToUserModel{
		UserID:     assignBook.UserID,
		BookID:     assignBook.BookID,
		AssignedAt: time.Now(),
	}
	err = ctrl.Repo.DB.Create(&newAssign).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"Status": http.StatusBadRequest, "message": "could not assign book to user"})
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": http.StatusOK, "message": "book assigned to user successfully"})
	return nil
}
func (ctrl *AssignBookToUserController) GetBookByUserId(c *fiber.Ctx) error {
	id := c.Params("userId")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"Status": http.StatusBadRequest, "message": "id must not be empty"})
		return nil
	}
	var user models.UserModel
	err := ctrl.Repo.DB.First(&user, "id=?", id).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "user not found"})
		return err
	}
	var assignedBook []models.AssignBookToUserModel
	err = ctrl.Repo.DB.Where("user_id=?", id).Find(&assignedBook).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "book not found"})
		return err
	}
	var response []map[string]interface{}
	for _, assignedBook := range assignedBook {
		var book models.BookModels

		err = ctrl.Repo.DB.Where("id = ?", assignedBook.BookID).Find(&book).Error
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"status": http.StatusBadRequest, "message": "book not found"})
			return nil
		}
		response = append(response, map[string]interface{}{
			"book":       book,
			"user":       user,
			"assignedAt": assignedBook.AssignedAt,
		})
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": http.StatusOK, "message": "request successfully", "data": response})

	return nil
}
func (ctrl *AssignBookToUserController) GetBookByBookId(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	if bookId == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "id must not be empty"})
		return nil
	}
	var book models.BookModels
	err := ctrl.Repo.DB.First(&book, "id=?", bookId).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "book not found"})
		return err
	}
	var assignUser []models.AssignBookToUserModel
	err = ctrl.Repo.DB.Where("book_id=?", bookId).Find(&assignUser).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "book not found"})
		return err
	}
	var response []map[string]interface{}
	for _, assignUser := range assignUser {
		fmt.Println("AssignUser.UserID:", assignUser.UserID)

		var user models.UserModel
		err = ctrl.Repo.DB.Where("id = ?", assignUser.UserID).Find(&user).Error
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"status": http.StatusBadRequest, "message": "user not found"})
			return nil
		}
		response = append(response, map[string]interface{}{
			"book":       book,
			"user":       user,
			"assignedAt": assignUser.AssignedAt,
		})
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": http.StatusOK, "message": "request successfully", "data": response})
	return nil
}
func (ctrl *AssignBookToUserController) ReturnBook(c *fiber.Ctx) error {

	var returnBook models.AssignBookToUserModel
	err := c.BodyParser(&returnBook)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "invalid request body"})
	}
	var assignedBook models.AssignBookToUserModel
	err = ctrl.Repo.DB.Where("BookID = ? AND UserID = ?", returnBook.BookID, returnBook.UserID).First(&assignedBook).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "book not found"})
		return err
	}
	assignedBook.ReturnAt = time.Now()
	err = ctrl.Repo.DB.Save(&assignedBook).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "could not return book"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": http.StatusOK, "message": "book returned successfully"})
	return nil
}
