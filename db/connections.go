package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	connStr := "host=localhost port=5432 user=postgres dbname=coinkeeper_db password=postgres"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Println("Connected to database")

	dbConn = db
	return nil
}

func CloseDBConn() error {
	//err := dbConn.Close()
	//if err != nil {
	//	return err
	//}

	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
