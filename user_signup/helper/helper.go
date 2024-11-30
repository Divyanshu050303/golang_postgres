package helper

import (
	"divyanshu050303/user_signup/models"
	"strings"

	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(user models.UserSignUpModels) (accessToken string, refreshToken string, err error) {
	accessTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessToken, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	refreshTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 168).Unix(),
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshToken, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return accessToken, refreshToken, err
}
func ValidateToken(tokenString string) (claims jwt.MapClaims, err error) {
	fmt.Println("token:", tokenString)
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	fmt.Println("token:", token)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("claims:", claims)
	fmt.Println("ok:", ok)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	return claims, nil
}
