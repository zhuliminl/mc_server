package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/controllers"
	"github.com/zhuliminl/mc_server/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	user := new(controllers.UserController)
	router.Use(middlewares.Auth())
	router.GET("/user", user.GetUser)
	router.PUT("/user", user.UpdateUser)

	return router
}
