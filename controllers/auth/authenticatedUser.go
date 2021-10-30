package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func GetAuthenticatedUser(c *gin.Context) {
	var db = config.GetDBInstance()
	var secret, _ = os.LookupEnv("SECRET_KEY")
	var secretByte = []byte(secret)
	tokenStr, cookieErr := c.Cookie("token")
	if cookieErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: "Token not found",
		})
		return
	}

	//cred := models.Credential{}
	token, tokenErr := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secretByte, nil
	})
	if tokenErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: tokenErr.Error(),
		})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: "error",
		})
	}

	userId, _ := claims["UserId"].(string)
	username, _ := claims["Username"].(string)

	cred := models.Credential{
		UserId:   userId,
		Username: username,
	}

	registeredUser := models.User{}

	whereMap := make(map[string]string)
	whereMap["id"] = cred.UserId
	selectStr := util.GenerateSelectFromTableString("USER", whereMap)

	if err := db.QueryRow(selectStr).Scan(&registeredUser.Id, &registeredUser.Username, &registeredUser.Email, &registeredUser.Password); err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.USER_NOT_FOUND,
		})
		return
	}
	c.JSON(200, util.ResponseUserEvent{
		Success: true,
		Message: models.UserResponse{
			Email:    registeredUser.Email,
			Username: registeredUser.Username,
		},
	})
}
