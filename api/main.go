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
	// "gorm.io/gorm"
)

var (
	userHandler *handlers.UserHandler
	roleHandler *handlers.RoleHandler
	projectHandler *handlers.ProjectHandler
	taskHandler *handlers.TaskHandler
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
	config.SyncDB()

	userService := services.NewUserService()
	roleService := services.NewRoleService()
	projectService := services.NewProjectService()
	taskService := services.NewTaskService()

	userHandler = handlers.NewUserHandler(userService)
	roleHandler = handlers.NewRoleHandler(roleService)
	projectHandler = handlers.NewProjectHandler(projectService)
	taskHandler = handlers.NewTaskHandler(taskService)

	initializeDefaultData(roleService, userService)
	// DatabaseMiddleware(config.DB)
}

// func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Set("db", db)
//         c.Next()
//     }
// }

func initializeDefaultData(roleService *services.RoleService, userService *services.UserService) {
	roles := []string{"Administrador", "Usuario"}
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
				roles.GET("/:roleId", roleHandler.GetRole)
				roles.PUT("/:roleId", roleHandler.UpdateRole)
				roles.DELETE("/:roleId", roleHandler.DeleteRole)
			}

			// User management (admin only)
			users := admin.Group("/users")
			{
				users.GET("/", userHandler.ListUsers)
				users.GET("/:userId", userHandler.GetUser)
				users.PUT("/:userId", userHandler.UpdateUser)
				users.DELETE("/:userId", userHandler.DeleteUser)
				users.POST("/:userId/:roleId", userHandler.AddRoleToUser)
				users.DELETE("/:userId/:roleId", userHandler.RemoveRoleFromUser)
			}

			// Project management
			projects := admin.Group("/projects")
			{
				projects.GET("/", projectHandler.ListProjects)
				projects.POST("/", projectHandler.CreateProject)
				projects.GET("/:projectId", projectHandler.GetProjectById)
				projects.PUT("/:projectId", projectHandler.UpdateProject)
				projects.DELETE("/:projectId", projectHandler.DeleteProject)
				projects.GET("/user/:userId", projectHandler.ListProjectsByUserId)	
				projects.POST("/:projectId/user/:userId", projectHandler.AddUserToProject)
				projects.DELETE("/:projectId/user/:userId", projectHandler.RemoveUserFromProject)
				projects.POST("/:projectId/task/:taskId", projectHandler.AddTaskToProject)
				projects.DELETE("/:projectId/task/:taskId", projectHandler.RemoveTaskFromProject)
			}

			tasks := admin.Group("/tasks")
			{
				tasks.GET("/", taskHandler.ListTasks)
				tasks.POST("/", taskHandler.CreateTask)
				tasks.GET("/:taskId", taskHandler.GetTaskById)
				tasks.PUT("/:taskId", taskHandler.UpdateTask)
				tasks.DELETE("/:taskId", taskHandler.DeleteTask)
			}
		}

		// Routes accessible by both Admin and Reader
		users := api.Group("/users")
		users.Use(middlewares.AuthMiddleware("Usuario"), middlewares.IsOwner("user"))
		{
			users.GET("/:userId" ,userHandler.GetUser)
			users.PUT("/:userId", userHandler.UpdateUser)
			users.DELETE("/:userId", userHandler.DeleteUser)
		}

		projects := api.Group("/projects")
		projects.Use(middlewares.AuthMiddleware("Usuario"))
		{
			projects.POST("/", projectHandler.CreateProject)
			projects.GET("/user/:userId",middlewares.IsOwner("user"), projectHandler.ListProjectsByUserId)
			projects.Use(middlewares.IsOwner("project"))
			{
				projects.GET("/:projectId", projectHandler.GetProjectById)
				projects.PUT("/:projectId", projectHandler.UpdateProject)
				projects.DELETE("/:projectId", projectHandler.DeleteProject)
				projects.POST("/:projectId/user/:userId", projectHandler.AddUserToProject)
				projects.DELETE("/:projectId/user/:userId", projectHandler.RemoveUserFromProject)
				projects.POST("/:projectId/task/:taskId", projectHandler.AddTaskToProject)
				projects.DELETE("/:projectId/task/:taskId", projectHandler.RemoveTaskFromProject)
			}
			
		}

		tasks := api.Group("/tasks")
		tasks.Use(middlewares.AuthMiddleware("Usuario"))
		{
			tasks.POST("/", taskHandler.CreateTask)
			tasks.Use(middlewares.IsOwner("task"))
			{
				tasks.GET("/:taskId", taskHandler.GetTaskById)
				tasks.PUT("/:taskId", taskHandler.UpdateTask)
				tasks.DELETE("/:taskId", taskHandler.DeleteTask)
			}

		}
	}
}