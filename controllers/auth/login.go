package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
	"golang.org/x/crypto/bcrypt"
)

func Login(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	rw.Header().Set("Content-Type", "application/json")
	var user models.LoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	credentialUser := models.User{}
	selectStr := fmt.Sprintf(`SELECT * FROM public."USER" WHERE email = '%s'`, user.Email)

	if scanErr := db.QueryRow(selectStr).Scan(&credentialUser.Id, &credentialUser.Username, &credentialUser.Email, &credentialUser.Password); scanErr != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	if util.IsEmpty(credentialUser.Email) {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if hashErr := bcrypt.CompareHashAndPassword([]byte(credentialUser.Password), []byte(user.Password)); hashErr != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	if tokenStr, authErr := credentialUser.AuthenticateUser(credentialUser.Id); authErr != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		expirationTime := time.Now().Add(time.Hour * 120)
		http.SetCookie(rw, &http.Cookie{
			Name:    "token",
			Value:   tokenStr,
			Expires: expirationTime,
		})
	}

}
