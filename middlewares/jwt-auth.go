package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zhuliminl/mc_server/controllers"
	"github.com/zhuliminl/mc_server/service"
)

func JWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("token")

		if tokenHeader == "" {
			response := controllers.BuildErrorResponse("token 参数不存在", 401, "请携带的 token 信息", controllers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := jwtService.ValidateToken(tokenHeader)
		if err != nil {
			response := controllers.BuildErrorResponse("token 校验失败", 401, err.Error(), controllers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			userId := claims["userId"]
			if userId == "" {
				response := controllers.BuildErrorResponse("token 校验失败, token 信息中用户信息为空", 401, "", controllers.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			c.Set("CurrentUserId", userId)

			log.Println("saul >>>>>>>>>>>>>>>>>>>>>>>>>>>> Claims: ", userId)

			c.Next()
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			response := controllers.BuildErrorResponse("token 校验失败, That's not even a token", 401, err.Error(), controllers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet
			response := controllers.BuildErrorResponse("token 校验失败, Timing is everything", 401, err.Error(), controllers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		} else {
			response := controllers.BuildErrorResponse("token 校验失败, Couldn't handle this toke", 401, err.Error(), controllers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
