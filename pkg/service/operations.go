package service

import (
	"coinkeeper/models"
	"coinkeeper/pkg/repository"
)

func GetAllOperations(userID uint, query string) (operations []models.Operation, err error) {
	operations, err = repository.GetAllOperations(userID, query)
	if err != nil {
		return nil, err
	}

	return operations, nil
}

func GetOperationByID(userID, operationID uint) (o models.Operation, err error) {
	o, err = repository.GetOperationByID(userID, operationID)
	if err != nil {
		return models.Operation{}, err
	}

	return o, nil
}

func CreateOperation(o models.Operation) error {
	if err := repository.CreateOperation(o); err != nil {
		return err
	}

	return nil
}

func UpdateOperation(o models.Operation) error {
	if err := repository.UpdateOperation(o); err != nil {
		return err
	}

	return nil
}
