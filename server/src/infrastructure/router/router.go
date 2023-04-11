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
	userRoute(e, db)
}

func userRoute(e *echo.Echo, db *gorm.DB) {
	userTable := persistence.NewUserTable(db)
	userFactory := factory.NewUserFactory()
	userDomainService := domain_service.NewUserService(userTable, userFactory)
	userService := service.NewUserService(userDomainService)
	userPresenter := presenter.NewUserPresenter()
	userHandler := handler.NewUserHandler(userService, userPresenter)

	e.GET("/users/:id", userHandler.GetUser)
	e.GET("/users", userHandler.GetUsers)
	e.POST("/users", userHandler.PostUser)
}
