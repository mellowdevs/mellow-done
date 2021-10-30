package models

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Credential struct {
	Username string
	UserId   string
	jwt.StandardClaims
}

func (cred *Credential) AuthenticateLogin(user User) (string, error) {
	var secret, _ = os.LookupEnv("SECRET_KEY")
	cred.UserId = user.Id
	cred.Username = user.Username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cred)

	cred.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenStr, err
}
