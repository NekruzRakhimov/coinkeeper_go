package controllers

import (
	"coinkeeper/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/ping", PingPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	userG := router.Group("/users")
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id", UpdateUser)
		userG.DELETE("/:id")
		userG.PATCH("/:id")
	}

	/*
		/operations -> checkUserAuthentication -> GetAllOperations
	*/
	operationsG := router.Group("/operations", checkUserAuthentication)
	{
		operationsG.GET("", GetAllOperations)
		operationsG.POST("", CreateOperation)
		operationsG.GET("/:id", GetOperationByID)
		operationsG.PUT("/:id", UpdateOperation) // admin permitted only
		operationsG.DELETE("/:id")
		operationsG.PATCH("/:id")
	}

	operationTypesG := router.Group("/operation-types")
	{
		operationTypesG.GET("")
		operationTypesG.POST("")
		operationTypesG.GET("/:id")
		operationTypesG.PUT("/:id")
		operationTypesG.DELETE("/:id")
	}

	operationCategoriesG := router.Group("/operation-categories")
	{
		operationCategoriesG.GET("")
		operationCategoriesG.POST("")
		operationCategoriesG.GET("/:id")
		operationCategoriesG.PUT("/:id")
		operationCategoriesG.DELETE("/:id")
	}

	return router
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
