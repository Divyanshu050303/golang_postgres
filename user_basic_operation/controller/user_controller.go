package controller

import (
	"github.com/divyanshu050303/user_basic_operation/repository"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Repo *repository.UserRepository
}

func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	return nil
}
func (ctrl *UserController) GetUsers(c *fiber.Ctx) error {
	return nil
}
func (ctrl *UserController) GetUserById(c *fiber.Ctx) error {
	return nil
}
func (ctrl *UserController) DeleteUser(c *fiber.Ctx) error {
	return nil
}
func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	return nil
}
