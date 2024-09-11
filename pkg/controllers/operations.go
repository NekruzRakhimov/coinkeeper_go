package controllers

import (
	"coinkeeper/errs"
	"coinkeeper/models"
	"coinkeeper/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllOperations
// @Summary Get All Operations
// @Security ApiKeyAuth
// @Tags operations
// @Description get list of all operations
// @ID get-all-operations
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.Operation
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/operations [get]
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

// GetOperationByID
// @Summary Get Operation By ID
// @Security ApiKeyAuth
// @Tags operations
// @Description get operation by ID
// @ID get-operation-by-id
// @Produce json
// @Param id path integer true "id of the operation"
// @Success 200 {object} models.Operation
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/operations/{id} [get]
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

// CreateOperation
// @Summary Create Operation
// @Security ApiKeyAuth
// @Tags operations
// @Description create new operation
// @ID create-operation
// @Accept json
// @Produce json
// @Param input body models.Operation true "new operation info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/operations [post]
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

	c.JSON(http.StatusCreated, defaultResponse{Message: "operation created successfully"})

}

// UpdateOperation
// @Summary Update Operation
// @Security ApiKeyAuth
// @Tags operations
// @Description update existed operation
// @ID update-operation
// @Accept json
// @Produce json
// @Param id path integer true "id of the operation"
// @Param input body models.Operation true "operation update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/operations/{id} [put]
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

	c.JSON(http.StatusOK, defaultResponse{Message: "operation updated successfully"})
}

// DeleteOperation
// @Summary Delete Operation By ID
// @Security ApiKeyAuth
// @Tags operations
// @Description delete operation by ID
// @ID delete-operation-by-id
// @Param id path integer true "id of the operation"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/operations/{id} [delete]
func DeleteOperation(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	operationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err = service.DeleteOperation(operationID, userID); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, defaultResponse{Message: "operation deleted successfully"})
}
