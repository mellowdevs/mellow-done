package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/mellowdevs/mellow-done/models"
)

func GetUser(rw http.ResponseWriter, r *http.Request) models.Credential {
	user := models.Credential{}
	var secretKey, _ = os.LookupEnv("SECRET_KEY")
	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			rw.WriteHeader(http.StatusUnauthorized)
			return user
		}
		rw.WriteHeader(http.StatusBadRequest)
		return user
	}
	tokenStr := token.Value
	claims := &models.Credential{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) { return []byte(secretKey), nil })
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			rw.WriteHeader(http.StatusUnauthorized)
			return user
		}
		rw.WriteHeader(http.StatusBadRequest)
		return user
	}
	if !tkn.Valid {
		rw.WriteHeader(http.StatusUnauthorized)
		return user
	}

	rw.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
	return *claims
}
