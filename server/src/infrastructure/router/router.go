package router

import (
	"github.com/labstack/echo/v4"
	"github.com/teshimafu/lazyPM/server/src/domain/factory"
	domain_service "github.com/teshimafu/lazyPM/server/src/domain/service"
	"github.com/teshimafu/lazyPM/server/src/infrastructure/auth"
	"github.com/teshimafu/lazyPM/server/src/infrastructure/persistence"
	"github.com/teshimafu/lazyPM/server/src/interfaces/handler"
	"github.com/teshimafu/lazyPM/server/src/interfaces/middleware"
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
	userService := service.NewUserService(userDomainService)
	authService := service.NewAuthService(userDomainService, authDomainService)

	// presenter
	userPresenter := presenter.NewUserPresenter()

	// handler
	authMiddleware := middleware.NewAuthMiddleware(authService)
	userHandler := handler.NewUserHandler(userService, userPresenter)
	authHandler := handler.NewAuthHandler(authService, userPresenter, userFactory)

	e.POST("/signup", authHandler.PostSignup)
	e.POST("/signin", authHandler.PostSignin)

	protected := e.Group("")
	protected.Use(authMiddleware.Middleware)
	protected.GET("/users", userHandler.GetUsers)
}
