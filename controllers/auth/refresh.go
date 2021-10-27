package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mellowdevs/mellow-done/models"
)

func Refresh(rw http.ResponseWriter, r *http.Request) {
	var secretKey, _ = os.LookupEnv("SECRET_KEY")
	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := token.Value
	claims := &models.Credential{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) { return []byte(secretKey), nil })
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 1*time.Second {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	expireTime := time.Now().Add(time.Hour * 120)
	claims.ExpiresAt = expireTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := newToken.SignedString([]byte(secretKey))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expireTime,
	})

}
