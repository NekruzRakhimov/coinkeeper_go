package repository

import (
	"coinkeeper/db"
	"coinkeeper/models"
	"fmt"
)

func AddOperation(o models.Operation) error {
	return nil
}

func GetAllOperations(userID uint) ([]models.Operation, error) {
	var operations []models.Operation
	err := db.GetDBConn().Model(&models.Operation{}).
		Joins("JOIN users ON users.id = operations.user_id").
		Where("operations.user_id = ?", userID).
		Order("operations.id").
		Find(&operations).Error
	if err != nil {
		return nil, err
	}
	return operations, nil
}

func GetOperationByID(userID, operationID uint) (o models.Operation, err error) {
	err = db.GetDBConn().Model(&models.Operation{}).
		Joins("JOIN users ON users.id = operations.user_id").
		Where("operations.user_id = ? AND operations.id = ?", userID, operationID).
		First(&o).Error
	if err != nil {
		return models.Operation{}, err
	}
	return o, nil
}

func GetTotalByOperationType(operationTypeID int) (total float64, err error) {
	err = db.GetDBConn().Raw(db.GetTotalAmountByOperationType, operationTypeID).Pluck("sum", &total).Error
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return total, nil
}

func UpdateOperation(o models.Operation) error {

	err := db.GetDBConn().Save(&o).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteOperation(id int) error {

	db.GetDBConn().Table("operations").Where("id = ?", id).Update("is_deleted", true)

	return nil
}
