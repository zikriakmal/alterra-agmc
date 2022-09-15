package middlewares

import (
	"day_2/config"
	"day_2/constants"
	"day_2/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

func BasicAuthDb(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User
	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) uint {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return uint(userId)
	}
	return 0
}
