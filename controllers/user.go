package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/service"
)


type UserController interface {
	Profile(context *gin.Context)
}

type userController struct{
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (ctl *userController) Profile (c *gin.Context) {
	id := "1"
	user := ctl.userService.Profile(id)
	c.JSON(http.StatusOK, user)
	c.Abort()
}



// func (u UserController) UpdateUser(c *gin.Context) {
// 	var user forms.UserSignUp
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "updateUser"})
// 	c.Abort()
// 	return
// }
