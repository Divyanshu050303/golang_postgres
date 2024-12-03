package controller

import (
	"divyanshu050303/user_signup/helper"
	"divyanshu050303/user_signup/models"
	"divyanshu050303/user_signup/repository"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	Repo *repository.UserRepository
}

func (ctrl *UserController) SignupUser(c *fiber.Ctx) error {
	userModel := models.UserSignUpModels{
		ID: uuid.New().String(),
	}
	err := c.BodyParser(&userModel)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": http.StatusBadRequest, "message": "invalid request body"})
		return err
	}
	var existingUser models.UserSignUpModels

	err = ctrl.Repo.DB.Where("user_email=?", userModel.UserEmail).Find(&existingUser).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  http.StatusBadRequest,
				"message": "invalid request body"})
		return err
	}
	if existingUser.ID != "" {
		c.Status(http.StatusConflict).JSON(
			&fiber.Map{
				"Status":  http.StatusConflict,
				"message": "user already exists"})
		return nil
	}
	err = ctrl.Repo.DB.Create(&userModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  http.StatusBadRequest,
				"message": "could not create the user"})
		return err
	}
	accessToken, refreshToken, err := helper.GenerateTokens(userModel)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "could not generate tokens"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": http.StatusOK,
			"message":       "user created successfully",
			"access_token":  accessToken,
			"refresh_token": refreshToken})
	return nil
}
func (ctrl *UserController) LoginUser(c *fiber.Ctx) error {
	var loginModles models.LoginModels
	err := c.BodyParser(&loginModles)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  http.StatusBadRequest,
				"message": "invalid request body"})
		return err
	}
	var user models.UserSignUpModels
	err = ctrl.Repo.DB.Where("user_email=? AND user_password=?", loginModles.UserEmail, loginModles.UserPassword).Find(&user).Error
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"status": http.StatusUnauthorized, "message": "invalid credentials"})
		return err
	}
	if user.ID == "" {
		c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"status": http.StatusUnauthorized, "message": "invalid credentials"})
		return nil
	}
	// Generate tokens
	accessToken, refreshToken, err := helper.GenerateTokens(user)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"status": http.StatusInternalServerError, "message": "could not generate tokens"})
		return err
	}

	// Return tokens
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":        http.StatusOK,
		"message":       "logged in successfully",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
	return nil
}
func (ctrl *UserController) GetUsers(c *fiber.Ctx) error {
	req := c.Request()
	tokenString := string(req.Header.Peek("Authorization"))
	fmt.Println("tokenString:", tokenString)
	if tokenString == "" {
		c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"status": http.StatusUnauthorized, "message": "token is missing"})
		return nil
	}

	_, err := helper.ValidateToken(tokenString)
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"status": http.StatusUnauthorized, "message": "invalid tokens"})
		return nil
	}
	var users []models.UserSignUpModels
	errosr := ctrl.Repo.DB.Find(&users).Error
	if errosr != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"status": http.StatusInternalServerError, "message": "could not retrieve users"})
		return err
	}
	c.Status(http.StatusOK).JSON(users)
	return nil

}
func (ctrl *UserController) ForgotPassword(c *fiber.Ctx) error {
	return nil
}
func (ctrl *UserController) ResetPassword(c *fiber.Ctx) error {

	return nil
}
