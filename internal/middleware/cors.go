package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CorsMiddleware() echo.MiddlewareFunc {
	return middleware.CORS()
}
