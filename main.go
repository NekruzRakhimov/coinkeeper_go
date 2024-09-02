package main

import (
	"coinkeeper/configs"
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/pkg/controllers"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(errors.New(fmt.Sprintf("error loading .env file. Error is %s", err)))
	}

	err = configs.ReadSettings()
	if err != nil {
		panic(err)
	}

	err = logger.Init()
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

	//MainServer := new(server.Server)
	//go func() {
	//	if err := MainServer.Run(configs.AppSettings.AppParams.PortRun, handlers.InitRoutes()); err != nil {
	//		log.Fatalf("Error while running http server. Error is %s", err.Error())
	//	}
	//}()
	//fmt.Println("TodoApp Started its work")
	//fmt.Printf("Server is listening port: %s\n", configs.AppSettings.AppParams.PortRun)
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	//<-quit
	///**********************************************************/
	//
	///***************** Shutting App Down *****************/
	//fmt.Println("TodoApp Shutting Down")
	//if err := MainServer.Shutdown(context.Background()); err != nil {
	//	log.Fatalf("error while shutting server down. Error is: %s", err.Error())
	//}
	//
	//if err := db.GetDBConn().Close(); err != nil {
	//	log.Fatalf("error while closing DB. Error is: %s", err.Error())
	//}

	err = controllers.RunRoutes()
	if err != nil {
		return
	}
}
