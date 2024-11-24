package controller

import (
	"net/http"

	"github.com/divyanshu050303/user_basic_operation/models"
	"github.com/divyanshu050303/user_basic_operation/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	Repo *repository.UserRepository
}

func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	user := models.UserModel{
		ID: uuid.New().String(),
	}
	err := c.BodyParser(&user)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}
	err = ctrl.Repo.DB.Create(&user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"messge": "Could not create the user"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"status": http.StatusAccepted, "message": "User has been created successfully"})
	return nil
}
func (ctrl *UserController) GetUsers(c *fiber.Ctx) error {
	user := &[]models.UserModel{}
	err := ctrl.Repo.DB.Find(user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": http.StatusBadRequest, "message": "Counld not get book"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"status": http.StatusOK, "message": "User fetched successfully", "data": user})
	return nil
}
func (ctrl *UserController) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &models.UserModel{}
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": http.StatusBadRequest, "message": "id connot be empty"})
		return nil
	}
	err := ctrl.Repo.DB.Where("id=?", id).Find(user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": http.StatusBadRequest, "message": "Counld not get user"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"status": http.StatusOK, "message": "User fetched successfully", "data": user})
	return nil
}
func (ctrl *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &models.UserModel{}
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": http.StatusBadRequest, "message": "id connot be empty"})
		return nil
	}
	err := ctrl.Repo.DB.Delete(user, "id = ?", id).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": http.StatusBadRequest, "message": "Counld not delete user"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"status": http.StatusOK, "message": "User deleted successfully"})
	return nil
}
func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &models.UserModel{}
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": http.StatusBadRequest, "message": "id connot be empty"})
		return nil
	}
	err := ctrl.Repo.DB.First(&user, "id = ?", id).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": http.StatusBadRequest, "message": "Could Not update the user"})
		return err
	}
	var updateUser models.UserModel
	err = c.BodyParser(&updateUser)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": http.StatusBadRequest, "message": "invalid request body"})
		return err
	}
	user.UserName = updateUser.UserName
	user.UserEmail = updateUser.UserEmail
	user.UserPassword = updateUser.UserPassword
	err = ctrl.Repo.DB.Save(&user).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": http.StatusBadRequest, "message": "Could Not update the user"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"status": http.StatusOK, "message": "User updated successfully"})
	return nil
}
