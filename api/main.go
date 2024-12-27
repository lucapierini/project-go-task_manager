package main

import (
	"github.com/lucapierini/project-go-task_manager/config"
	"github.com/lucapierini/project-go-task_manager/middlewares"
	"github.com/lucapierini/project-go-task_manager/handlers"
	"github.com/gin-gonic/gin"
)

var (
	userHandler *handlers.UserHandler
	router *gin.Engine
)

func init(){
	config.LoadEnvVariables()
	config.ConnectDB()
	config.SyncDB()
}

func loadRoutes() {
	router = gin.Default()

	public := router.Group("/api")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
	
	}

	// Middleware for all routes - CORS
	router.Use(middlewares.CORSMiddleware())
	// Group with middleware CheckToken
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userHandler.ListUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	

}

func main() {

}