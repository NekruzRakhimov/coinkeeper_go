package repository

import (
	"coinkeeper/db"
	"coinkeeper/logger"
	"coinkeeper/models"
)

func CreateOperation(o models.Operation) error {
	err := db.GetDBConn().Create(&o).Error
	if err != nil {
		logger.Error.Println("[repository.CreateOperation] cannot create operation. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func GetAllOperations(userID uint, query string) ([]models.Operation, error) {
	var operations []models.Operation

	query = "%" + query + "%"

	err := db.GetDBConn().Model(&models.Operation{}).
		Joins("JOIN users ON users.id = operations.user_id").
		Where("operations.user_id = ? AND description iLIKE ?", userID, query).
		Order("operations.id").
		Find(&operations).Error
	if err != nil {
		logger.Error.Println("[repository.GetAllOperations] cannot get all operations. Error is:", err.Error())
		return nil, translateError(err)
	}
	return operations, nil
}

func GetOperationByID(userID, operationID uint) (o models.Operation, err error) {
	err = db.GetDBConn().Model(&models.Operation{}).
		Joins("JOIN users ON users.id = operations.user_id").
		Where("operations.user_id = ? AND operations.id = ?", userID, operationID).
		First(&o).Error
	if err != nil {
		logger.Error.Println("[repository.GetOperationByID] cannot get operation by id. Error is:", err.Error())
		return models.Operation{}, translateError(err)
	}
	return o, nil
}

func UpdateOperation(o models.Operation) error {
	err := db.GetDBConn().Save(&o).Error
	if err != nil {
		logger.Error.Println("[repository.UpdateOperation] cannot update operation. Error is:", err.Error())
		return translateError(err)
	}

	return nil
}

func DeleteOperation(operationID int, userID uint) error {
	err := db.GetDBConn().
		Table("operations").
		Where("id = ? AND user_id = ?", operationID, userID).
		Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteOperation] cannot delete operation. Error is:", err.Error())
		return err
	}

	return nil
}
