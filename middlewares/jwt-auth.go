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
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("saul >>>>>> . Claim[userId]: ", claims["userId"])
			log.Println("saul >>>>>>>>>>>> Claim[issuer] :", claims["issuer"])

			c.Next()
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			response := controllers.BuildErrorResponse("token 校验失败, That's not even a token", 401, err.Error(), controllers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet

			response := controllers.BuildErrorResponse("token 校验失败, Timing is everything", 401, err.Error(), controllers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		} else {
			response := controllers.BuildErrorResponse("token 校验失败, Couldn't handle this toke", 401, err.Error(), controllers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

		/*
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[userId]: ", claims["userId"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println("saul JWT ===========>>> ", err)
			response := controllers.BuildErrorResponse("token is not valid", 401, "未授权", controllers.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		*/
	}
}
