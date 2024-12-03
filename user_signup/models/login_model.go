package models

type LoginModels struct {
	UserEmail    *string `json:"userEmail"`
	UserPassword *string `json:"userPassword" validate:"min=8"`
}
