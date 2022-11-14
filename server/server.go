package server

import (
	"database/sql"

	"github.com/go-redis/redis/v9"
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var (
		db               *sql.DB                      = database.GetDB()
		userRepository   repository.UserRepository    = repository.NewUserRepository(db)
		userService      service.UserService          = service.NewUserService(userRepository)
		authService      service.AuthService          = service.NewAuthService(userRepository, userService)
		jwtService       service.JWTService           = service.NewJWTService()
		wechatService    service.WechatService        = service.NewWechatService(userRepository, userService, rdb)
		userController   controllers.UserController   = controllers.NewUserController(userService)
		authController   controllers.AuthController   = controllers.NewAuthController(authService, userService, jwtService)
		wechatController controllers.WechatController = controllers.NewWechatController(wechatService)
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
	router.POST("/loginByEmail", authController.LoginByEmail)
	router.POST("/loginByPhone", authController.LoginByPhone)
	router.POST("/registerByEmail", authController.RegisterByEmail)
	router.POST("/registerByPhone", authController.RegisterByPhone)

	router.POST("/openId", wechatController.GetOpenID)
	router.GET("/getMiniProgramLink", wechatController.GenerateAppLink)
	router.GET("/getMiniLinkStatus", wechatController.GetMiniLinkStatus)
	router.POST("/miniProgramScanOver", wechatController.ScanOver)
	router.POST("/loginWithEncryptedPhoneData", wechatController.LoginWithEncryptedPhoneData)

	router.Run(address + ":" + port)
}
