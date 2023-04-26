package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/teshimafu/lazyPM/server/src/usecase/service"
)

type AuthMiddleware struct {
	authService *service.AuthService
}

func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (a *AuthMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header format")
		}

		userID, err := a.authService.GetUserID(parts[1])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
		c.Set("user_id", userID)
		return next(c)
	}
}
