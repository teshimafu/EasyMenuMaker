package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teshimafu/lazyPM/server/migrations"
	"github.com/teshimafu/lazyPM/server/src/infrastructure/router"
)

func main() {
	db, err := gorm.Open(sqlite.Open("lazyPM.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = migrations.Migrate(db)
	if err != nil {
		panic(err)
	}

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
