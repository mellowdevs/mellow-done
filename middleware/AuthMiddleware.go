package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/util"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var secret, _ = os.LookupEnv("SECRET_KEY")
		var secretByte = []byte(secret)
		tokenStr, cookieErr := c.Cookie("token")
		if cookieErr != nil {
			c.JSON(400, util.ResponseMessage{
				Success: false,
				Message: util.NOT_AUTHED,
			})
			return
		}

		token, tokenErr := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return secretByte, nil
		})
		if tokenErr != nil {
			c.Abort()
			c.JSON(401, util.ResponseMessage{
				Success: false,
				Message: tokenErr.Error(),
			})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["UserId"].(string))
			c.Next()
		} else {
			c.Abort()
			c.JSON(400, util.ResponseMessage{
				Success: false,
				Message: "error",
			})
		}
	}

}
