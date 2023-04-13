package router

import (
	"github.com/labstack/echo/v4"
	"github.com/teshimafu/lazyPM/server/src/domain/factory"
	domain_service "github.com/teshimafu/lazyPM/server/src/domain/service"
	"github.com/teshimafu/lazyPM/server/src/infrastructure/persistence"
	"github.com/teshimafu/lazyPM/server/src/interfaces/handler"
	"github.com/teshimafu/lazyPM/server/src/interfaces/presenter"
	"github.com/teshimafu/lazyPM/server/src/usecase/service"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {
	// repository
	userTable := persistence.NewUserTable(db)

	// factory
	userFactory := factory.NewUserFactory()

	// domain service
	userDomainService := domain_service.NewUserService(userTable, userFactory)

	// application service
	userService := service.NewUserService(userDomainService)

	// presenter
	userPresenter := presenter.NewUserPresenter()

	// handler
	userHandler := handler.NewUserHandler(userService, userPresenter)
	authHandler := handler.NewAuthHandler(userService, userPresenter)

	e.GET("/users/:id", userHandler.GetUser)
	e.GET("/users", userHandler.GetUsers)

	e.POST("/signup", authHandler.PostSignup)
}
