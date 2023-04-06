package main

import (
	"github.com/teshimafu/lazyPM/src/cmd"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/teshimafu/lazyPM/migrations"
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
	cmd.Start(db)
}
