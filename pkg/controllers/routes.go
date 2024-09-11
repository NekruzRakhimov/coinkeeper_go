package controllers

import (
	"coinkeeper/configs"
	_ "coinkeeper/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", PingPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	apiG := router.Group("/api", checkUserAuthentication)

	userG := apiG.Group("/users")
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id", UpdateUser)
		userG.DELETE("/:id")
		userG.PATCH("/:id")
	}

	operationsG := apiG.Group("/operations")
	{
		operationsG.GET("", GetAllOperations)
		operationsG.POST("", CreateOperation)
		operationsG.GET("/:id", GetOperationByID)
		operationsG.PUT("/:id", UpdateOperation) // admin permitted only
		operationsG.DELETE("/:id", DeleteOperation)
	}

	operationTypesG := apiG.Group("/operation-types")
	{
		operationTypesG.GET("")
		operationTypesG.POST("")
		operationTypesG.GET("/:id")
		operationTypesG.PUT("/:id")
		operationTypesG.DELETE("/:id")
	}

	operationCategoriesG := apiG.Group("/operation-categories")
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
