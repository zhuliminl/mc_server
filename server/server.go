package server

import (
	"database/sql"
	"github.com/zhuliminl/mc_server/middlewares"

	"github.com/gin-gonic/gin"

	"github.com/zhuliminl/mc_server/config"
	"github.com/zhuliminl/mc_server/controllers"
	"github.com/zhuliminl/mc_server/database"
	"github.com/zhuliminl/mc_server/repository"
	"github.com/zhuliminl/mc_server/service"
)

func init() {
}

func StartServer() {
	var (
		db             *sql.DB                    = database.GetDB()
		userRepository repository.UserRepository  = repository.NewUserRepository(db)
		userService    service.UserService        = service.NewUserService(userRepository)
		authService    service.AuthService        = service.NewAuthService(userRepository)
		jwtService     service.JWTService         = service.NewJWTService()
		userController controllers.UserController = controllers.NewUserController(userService)
		authController controllers.AuthController = controllers.NewAuthController(authService, userService, jwtService)
	)
	c := config.GetConfig()
	address := c.GetString("server.address")
	port := c.GetString("server.port")

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/generateUser", userController.GenerateUsers)
	router.GET("/user", middlewares.JWT(jwtService), userController.GetByUserId)
	router.GET("/userAll", middlewares.JWT(jwtService), userController.GetAll)
	router.POST("/user", userController.Create)
	router.DELETE("/user", userController.DeleteByUserId)
	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)
	router.POST("/getToken", authController.GetToken)
	router.GET("/validateToken", authController.VerifyToken)

	router.Run(address + ":" + port)
}
