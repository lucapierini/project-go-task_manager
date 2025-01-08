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
	projectHandler *handlers.ProjectHandler
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
	config.SyncDB()

	userService := services.NewUserService()
	roleService := services.NewRoleService()
	projectService := services.NewProjectService()

	userHandler = handlers.NewUserHandler(userService)
	roleHandler = handlers.NewRoleHandler(roleService)
	projectHandler = handlers.NewProjectHandler(projectService)

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
				users.GET("/:id", userHandler.GetUser)
				users.PUT("/:id", userHandler.UpdateUser)
				users.DELETE("/:id", userHandler.DeleteUser)
				users.POST("/:id_user/:id_role", userHandler.AddRoleToUser)
				users.DELETE("/:id/:id_role", userHandler.RemoveRoleFromUser)
			}

			// Project management
			projects := admin.Group("/projects")
			{
				projects.GET("/", projectHandler.ListProjects)
				projects.POST("/", projectHandler.CreateProject)
				projects.GET("/:id", projectHandler.GetProjectById)
				projects.PUT("/:id", projectHandler.UpdateProject)
				projects.DELETE("/:id", projectHandler.DeleteProject)
				projects.GET("/user/:id", projectHandler.ListProjectsByUserId)	
				projects.POST("/:id/user/:userId", projectHandler.AddUsersToProject)
				projects.DELETE("/:id/user/:userId", projectHandler.RemoveUsersFromProject)
				projects.POST("/:id/task", projectHandler.AddTasksToProject)
				projects.DELETE("/:id/task/:taskId", projectHandler.RemoveTasksFromProject)
			}
		}

		// Routes accessible by both Admin and Reader
		users := api.Group("/users")
		users.Use(middlewares.AuthMiddleware("Administrador", "Lector"))
		{
			users.GET("/:id", userHandler.GetUser)
		}

		projects := api.Group("/projects")
		projects.Use(middlewares.AuthMiddleware("Administrador", "Usuario"))
		{
			projects.POST("/", projectHandler.CreateProject)
			
		}
	}
}