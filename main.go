package main

import (
	"coinkeeper/configs"
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/pkg/controllers"
	"coinkeeper/server"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
Vasya - 1111
Sadam - 2222

account(1111) - 100c
account(2222) + 100c
*/

func main() {
	// Загрузка .env файла
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %s", err)
	}

	// Чтение настроек
	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	// Инициализация логгера
	if err := logger.Init(); err != nil {
		log.Fatalf("Ошибка инициализации логгера: %s", err)
	}

	// Подключение к базе данных
	var err error
	err = db.ConnectToDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %s", err)
	}

	// Миграция базы данных
	if err = db.Migrate(); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %s", err)
	}

	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
			log.Fatalf("Ошибка при запуске HTTP сервера: %s", err)
		}
	}()

	// Ожидание сигнала для завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("Coinkeeper завершает работу")

	// Закрытие соединения с БД, если необходимо
	if sqlDB, err := db.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Ошибка при закрытии соединения с БД: %s", err)
		}
	} else {
		log.Fatalf("Ошибка при получении *sql.DB из GORM: %s", err)
	}

	if err = mainServer.Shutdown(context.Background()); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}
}
