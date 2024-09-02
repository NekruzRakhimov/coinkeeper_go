package controllers

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllOperations(c *gin.Context) {
	query := c.Query("q")

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	operations, err := service.GetAllOperations(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"operations": operations,
	})
}

func GetOperationByID(c *gin.Context) {
	userID := 8
	operationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	operation, err := service.GetOperationByID(uint(userID), uint(operationID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, operation)
}

func CreateOperation(c *gin.Context) {
	var o models.Operation
	if err := c.BindJSON(&o); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	o.UserID = userID

	if err := service.CreateOperation(o); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Operation created",
	})

}

func UpdateOperation(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if userRole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	operationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var o models.Operation
	if err = c.BindJSON(&o); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	o.ID = uint(operationID)
	o.UserID = userID

	if err = service.UpdateOperation(o); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Operation updated",
	})

}

func DeleteOperation(c *gin.Context) {

}
