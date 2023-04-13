package router

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/teshimafu/lazyPM/server/src/domain/factory"
	domain_service "github.com/teshimafu/lazyPM/server/src/domain/service"
	"github.com/teshimafu/lazyPM/server/src/infrastructure/auth"
	"github.com/teshimafu/lazyPM/server/src/infrastructure/persistence"
	"github.com/teshimafu/lazyPM/server/src/interfaces/handler"
	"github.com/teshimafu/lazyPM/server/src/interfaces/presenter"
	"github.com/teshimafu/lazyPM/server/src/usecase/service"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {
	// repository
	userTable := persistence.NewUserTable(db)
	tokenGenerator := auth.NewJWTGenerator()

	// factory
	userFactory := factory.NewUserFactory(userTable)

	// domain service
	userDomainService := domain_service.NewUserService(userTable)
	authDomainService := domain_service.NewAuthService(tokenGenerator)

	// application service
	userService := service.NewUserService(userDomainService, authDomainService)

	// presenter
	userPresenter := presenter.NewUserPresenter()

	// handler
	userHandler := handler.NewUserHandler(userService, userPresenter)
	authHandler := handler.NewAuthHandler(userService, userPresenter, userFactory)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Path(), "/signup") || strings.HasPrefix(c.Path(), "/signin") {
				return next(c)
			}
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header format")
			}

			token := parts[1]
			c.Set("token", token)

			return next(c)
		}
	})

	e.POST("/signup", authHandler.PostSignup)
	e.POST("/signin", authHandler.PostSignin)

	e.GET("/users", userHandler.GetUsers)
}
