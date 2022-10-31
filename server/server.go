package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/zhuliminl/mc_server/config"
	"github.com/zhuliminl/mc_server/controllers"
	"github.com/zhuliminl/mc_server/database"
	"github.com/zhuliminl/mc_server/repository"
	"github.com/zhuliminl/mc_server/service"
)

var (
	db             *sql.DB                    = database.ConnectDB()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	userService    service.UserService        = service.NewUserService(userRepository)
	userController controllers.UserController = controllers.NewUserController(userService)
)

func StartServer() {
	c := config.GetConfig()
	address := c.GetString("server.address")
	port := c.GetString("server.port")

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/user", userController.Profile)

	router.Run(address + ":" + port)
}
