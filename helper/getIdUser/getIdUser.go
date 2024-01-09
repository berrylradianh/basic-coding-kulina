package getiduser

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetIdUser(c echo.Context) (string, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	claimsID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("Invalid user ID format")
	}

	return claimsID, nil
}
