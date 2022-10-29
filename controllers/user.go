package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/forms"
)

type UserController struct{}

func (u UserController) GetUser(c *gin.Context) {

	var id forms.UserId
	err := c.BindQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ok"})
	c.Abort()
	return
}

func (u UserController) UpdateUser(c *gin.Context) {
	var user forms.UserSignUp
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updateUser"})
	c.Abort()
	return
}
