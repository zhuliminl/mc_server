package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("middlewraes => 中间件")
		c.Next()
	}
}
