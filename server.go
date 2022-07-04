package main

import (
	"gin-api/config"
	"gin-api/controller"
	"gin-api/repository"
	"gin-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	apiRoutes := server.Group("/api")
	{
		apiRoutes.POST("/login", authController.Login)
		apiRoutes.POST("/register", authController.Register)
	}

	server.Run(":8080")
}
