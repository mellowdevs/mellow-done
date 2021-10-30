package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var db = config.GetDBInstance()
	userRequest := models.LoginUser{}
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Fatal(err)
	}

	whereMap := make(map[string]string)
	whereMap["email"] = userRequest.Email
	selectStr := util.GenerateSelectFromTableString("USER", whereMap)

	registeredUser := models.User{}

	if err := db.QueryRow(selectStr).Scan(&registeredUser.Id, &registeredUser.Username, &registeredUser.Email, &registeredUser.Password); err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.USER_NOT_FOUND,
		})
		return
	}

	if passwordErr := bcrypt.CompareHashAndPassword([]byte(registeredUser.Password), []byte(userRequest.Password)); passwordErr != nil {

		fmt.Println(registeredUser.Password)
		fmt.Println(userRequest.Password)

		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.INCORRECT_PASSWORD,
		})
		return
	}

	credential := models.Credential{}

	token, tokenErr := credential.AuthenticateLogin(registeredUser)
	if tokenErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: tokenErr.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("token", token, 2000, "/", "", false, true)
	c.JSON(200, util.ResponseUserEvent{
		Success: true,
		Message: models.UserResponse{
			registeredUser.Email,
			registeredUser.Username,
		},
	})

}
