package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middlewraes => 中间件")
		c.Next()
	}
}
