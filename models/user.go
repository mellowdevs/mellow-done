package models

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

type Credential struct {
	Username string
	jwt.StandardClaims
}

func (user *User) EncryptPassword(password string) {
	if passHash, err := bcrypt.GenerateFromPassword([]byte(password), 12); err != nil {
		log.Fatal(err)
	} else {
		user.Password = string(passHash)
	}

}

func (user *User) AuthenticateUser(id string) (string, error) {
	var secretKey, _ = os.LookupEnv("SECRET_KEY")
	expirationTime := time.Now().Add(time.Hour * 120)
	claims := &Credential{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if tokenString, err := token.SignedString([]byte(secretKey)); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}
