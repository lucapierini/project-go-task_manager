package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/project-go-task_manager/config"
	"github.com/lucapierini/project-go-task_manager/dto"
	"github.com/lucapierini/project-go-task_manager/handlers"
	"github.com/lucapierini/project-go-task_manager/middlewares"
	"github.com/lucapierini/project-go-task_manager/services"
)

var (
	userHandler *handlers.UserHandler
	roleHandler *handlers.RoleHandler
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
	config.SyncDB()

	userService := services.NewUserService()
	roleService := services.NewRoleService()

	userHandler = handlers.NewUserHandler(userService)
	roleHandler = handlers.NewRoleHandler(roleService)

	initializeDefaultData(roleService, userService)
}

func initializeDefaultData(roleService *services.RoleService, userService *services.UserService) {
	roles := []string{"Administrador", "Lector", "Usuario"}
	for _, role := range roles {
		if _, err := roleService.CreateRole(dto.RoleDto{Name: role}); err != nil {
			log.Printf("Error creating role %s: %v\n", role, err)
		}
	}

	adminUser := dto.UserDto{
		Username: "admin",
		Password: "admin",
		RoleIds:  []uint{1, 2},
		Email:    "admin@admin.com",
	}
	if _, err := userService.RegisterUser(adminUser); err != nil {
		log.Printf("Error creating admin user: %v\n", err)
	}
}

func main() {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	setupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Public routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			auth.POST("/refresh", handlers.RefreshTokenHandler)
		}

		// Protected routes
		admin := api.Group("/admin")
		admin.Use(middlewares.AuthMiddleware("Administrador"))
		{
			// Roles management
			roles := admin.Group("/roles")
			{
				roles.POST("/", roleHandler.CreateRole)
				roles.GET("/", roleHandler.ListRoles)
				roles.GET("/:id", roleHandler.GetRole)
				roles.PUT("/:id", roleHandler.UpdateRole)
				roles.DELETE("/:id", roleHandler.DeleteRole)
			}

			// User management (admin only)
			users := admin.Group("/users")
			{
				users.GET("/", userHandler.ListUsers)
				users.PUT("/:id", userHandler.UpdateUser)
				users.DELETE("/:id", userHandler.DeleteUser)
			}
		}

		// Routes accessible by both Admin and Reader
		users := api.Group("/users")
		users.Use(middlewares.AuthMiddleware("Administrador", "Lector"))
		{
			users.GET("/:id", userHandler.GetUser)
		}
	}
}