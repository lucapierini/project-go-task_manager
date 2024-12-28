package main

import (
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
	router *gin.Engine
)

func init(){
	config.LoadEnvVariables()
	config.ConnectDB()
	config.SyncDB()
	// config.CreateDefaultRoles()
	

	    // Initialize services
		userService := services.NewUserService()
		roleService := services.NewRoleService()
	
		// Initialize handlers with services
		userHandler = handlers.NewUserHandler(userService)
		roleHandler = handlers.NewRoleHandler(roleService)

		roleService.CreateRole(dto.RoleDto{Name: "Administrador"})
		roleService.CreateRole(dto.RoleDto{Name: "Lector"})
		roleService.CreateRole(dto.RoleDto{Name: "Usuario"})
		userService.RegisterUser(dto.UserDto{Username: "admin", Password: "admin", RoleIds: []uint{1, 2}, Email: "admin@admin.com"})
	
}

func main() {
	router = gin.Default()

	loadRoutes()

	router.Run(":8080")
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

	roleRoutes := router.Group("/roles")
	// Todas las rutas de roles están protegidas
	roleRoutes.Use(middlewares.CheckToken, middlewares.RoleMiddleware([]string{"Administrador"})) // Solo administradores pueden acceder
	{
		roleRoutes.POST("/", roleHandler.CreateRole)      // Crear rol
		roleRoutes.GET("/", roleHandler.ListRoles)        // Listar roles
		roleRoutes.GET("/:id", roleHandler.GetRole)       // Obtener rol por ID
		roleRoutes.PUT("/:id", roleHandler.UpdateRole)    // Actualizar rol
		roleRoutes.DELETE("/:id", roleHandler.DeleteRole)  // Eliminar rol
	}

	// Group with middleware CheckToken
	userRoutes := router.Group("/users")
	
	// Aplicar el middleware solo a las rutas específicas
	userRoutes.Use(middlewares.CheckToken, middlewares.RoleMiddleware([]string{"Lector", "Administrador"}))
	{
		userRoutes.GET("/", userHandler.ListUsers) // Esta ruta está protegida
		userRoutes.GET("/:id", userHandler.GetUser ) // Esta ruta también está protegida
	}

	userRoutes.Use(middlewares.RoleMiddleware([]string{"Administrador"}))
	{
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
