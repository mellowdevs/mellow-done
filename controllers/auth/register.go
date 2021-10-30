package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var db = config.GetDBInstance()
	user := models.User{}
	if bindErr := c.ShouldBindJSON(&user); bindErr != nil {
		log.Fatal(bindErr)
	}

	if util.IsEmpty(user.Username) {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: "Username " + util.EMPTY_FIELD_MSG,
		})
		return
	} else if len(user.Username) < 5 {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: "Username " + util.SHORT_FIELD_MSG,
		})
		return
	} else if util.IsEmpty(user.Email) {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: "Email " + util.EMPTY_FIELD_MSG,
		})
		return
	} else if util.IsEmailInvalid(user.Email) {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: "Email " + util.INVALID_FIELD_MSG,
		})
		return
	} else if util.IsEmpty(user.Password) {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: "Password " + util.EMPTY_FIELD_MSG,
		})
		return
	} else if len(user.Password) < 5 {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: "Password " + util.SHORT_FIELD_MSG,
		})
		return
	}

	whereEmailMap := make(map[string]string)
	whereEmailMap["email"] = user.Email
	whereUsernameMap := make(map[string]string)
	whereUsernameMap["username"] = user.Username
	selectEmailStr := util.GenerateSelectFromTableString("USER", whereEmailMap)
	selectUsernameStr := util.GenerateSelectFromTableString("USER", whereUsernameMap)

	registeredUser := models.User{}

	if err := db.QueryRow(selectEmailStr).Scan(&registeredUser.Id, &registeredUser.Username, &registeredUser.Email, &registeredUser.Password); err == nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.EMAIL_EXIST,
		})
		return
	}
	if err := db.QueryRow(selectUsernameStr).Scan(&registeredUser.Id, &registeredUser.Username, &registeredUser.Email, &registeredUser.Password); err == nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.USERNAME_EXIST,
		})
		return
	}

	hash, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if hashErr != nil {
		log.Fatal(hashErr)
	}
	user.Password = string(hash)

	fields := []string{"username", "email", "password"}
	values := []string{user.Username, user.Email, user.Password}
	insertStr := util.GenerateInsertIntoTableString("USER", fields, values)

	_, insertErr := db.Exec(insertStr)
	if insertErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: insertErr.Error(),
		})
		return
	}

	credential := models.Credential{}
	token, tokenErr := credential.AuthenticateLogin(user)
	if tokenErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: tokenErr.Error(),
		})
		return
	}
	secure := true
	if os.Getenv("GIN_ENV") == "development" {
		secure = false
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("token", token, 2000, "/", "", secure, true)
	c.JSON(200, util.ResponseUserEvent{
		Success: true,
		Message: models.UserResponse{
			Email:    user.Email,
			Username: user.Username,
		},
	})

}
