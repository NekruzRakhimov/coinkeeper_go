package db

import "coinkeeper/models"

func Migrate() error {
	err := dbConn.AutoMigrate(models.User{},
		models.OperationType{},
		models.OperationCategory{},
		models.Operation{})
	if err != nil {
		return err
	}
	return nil
}
