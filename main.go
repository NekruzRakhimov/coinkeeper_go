package main

import (
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/pkg/controllers"
)

func main() {
	err := logger.Init()
	if err != nil {
		panic(err)
	}

	err = db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	err = controllers.RunRoutes()
	if err != nil {
		return
	}
}
