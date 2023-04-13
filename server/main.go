package main

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teshimafu/lazyPM/server/migrations"
	"github.com/teshimafu/lazyPM/server/src/infrastructure/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	if key := os.Getenv("SECRET_KEY"); key == "" {
		panic("Error loading SECRET_KEY")
	}

	db, err := gorm.Open(sqlite.Open("lazyPM.db"), &gorm.Config{})
	if err != nil {
		panic("db access error")
	}
	err = migrations.Migrate(db)
	if err != nil {
		panic("db migrate error")
	}

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// initialize router
	router.Init(e, db)

	// start server
	log.Println("Server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
