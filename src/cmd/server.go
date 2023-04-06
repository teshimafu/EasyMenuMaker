package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/teshimafu/lazyPM/src/infrastructure/router"
)

func Start(db *gorm.DB) {

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// initialize router
	router.Init(e, db)

	// start server
	fmt.Println("Server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
