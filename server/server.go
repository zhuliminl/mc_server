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
		authService    service.AuthService        = service.NewAuthService(userRepository, userService)
		jwtService     service.JWTService         = service.NewJWTService()
		userController controllers.UserController = controllers.NewUserController(userService)
		authController controllers.AuthController = controllers.NewAuthController(authService, userService, jwtService)
	)

	JWTMiddleware := middlewares.JWT(jwtService)
	c := config.GetConfig()
	address := c.GetString("server.address")
	port := c.GetString("server.port")

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/generateUser", userController.GenerateUsers)
	router.GET("/user", JWTMiddleware, userController.GetByUserId)
	router.GET("/myInfo", JWTMiddleware, userController.GetMyInfo)
	router.GET("/userAll", JWTMiddleware, userController.GetAll)
	router.POST("/user", userController.Create)
	router.DELETE("/user", userController.DeleteByUserId)
	router.POST("/login", authController.Login)
	router.POST("/registerByEmail", authController.RegisterByEmail)
	router.POST("/registerByPhone", authController.RegisterByPhone)

	router.Run(address + ":" + port)
}
