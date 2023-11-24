package main

import (
	"sampleprj/config"
	"sampleprj/features/users/repository"
	"sampleprj/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	c := config.InitConfig()
	db, err := database.InitMySQL(*c)
	if err != nil {
		e.Logger.Fatal("cannot run database", err.Error())
	}
	err = db.AutoMigrate(repository.UserModel{})

	// redis := cache.InitRedis(*c)
	// if redis == nil {
	// 	e.Logger.Fatal("cannot run cache")
	// }

	// if err != nil {
	// 	e.Logger.Fatal("cannot run database")
	// }

	// mongo, err := database.InitMongoDB(*c)

	// if err != nil {
	// 	e.Logger.Fatal("cannot run mongodb", err.Error())
	// }

	e.Logger.Debug(db)
	// e.Logger.Debug(mongo)

	e.Logger.Fatal(e.Start(":8000"))
}
