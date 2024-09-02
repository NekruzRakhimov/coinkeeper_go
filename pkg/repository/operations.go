package repository

import (
	"coinkeeper/db"
	"coinkeeper/models"
)

func AddOperation(o models.Operation) error {
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
		return models.Operation{}, translateError(err)
	}
	return o, nil
}

func UpdateOperation(o models.Operation) error {

	err := db.GetDBConn().Save(&o).Error
	if err != nil {
		return translateError(err)
	}

	return nil
}

func DeleteOperation(id int) error {

	db.GetDBConn().Table("operations").Where("id = ?", id).Update("is_deleted", true)

	return nil
}
