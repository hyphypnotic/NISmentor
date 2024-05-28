package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is missing")
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims := &CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token signature is invalid")
			}
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
		}

		c.Set("username", claims.Username)

		return next(c)
	}
}
