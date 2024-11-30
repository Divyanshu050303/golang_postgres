package controller

import (
	"divyanshu050303/user_signup/repository"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Repo *repository.UserRepository
}

func (ctrl *UserController) SignupUser(c *fiber.Ctx) error {
	return nil
}
func (ctrl *UserController) LoginUser(c *fiber.Ctx) error {
	return nil
}
func (ctrl *UserController) GetUsers(c *fiber.Ctx) error {
	return nil
}
