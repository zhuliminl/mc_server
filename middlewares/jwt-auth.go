package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuliminl/mc_server/controllers"
	"github.com/zhuliminl/mc_server/service"
	"log"
	"net/http"
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
		log.Println("TOKEN", token)
		if err != nil {
			log.Println("JWT", err)

		}

		//if token.Valid {
		//	claims := token.Claims.(jwt.MapClaims)
		//	log.Println("Claim[userId]: ", claims["userId"])
		//	log.Println("Claim[issuer] :", claims["issuer"])
		//} else {
		//	log.Println("saul JWT ===========>>> ", err)
		//	response := controllers.BuildErrorResponse("token is not valid", 401, "未授权", controllers.EmptyObj{})
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		//}
	}
}
