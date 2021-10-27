package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func Register(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	rw.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.BadRequest(rw, util.PAYLOAD_ERR)
	} else {
		if userInvalid(&user) {
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			user.EncryptPassword(user.Password)
		}

		insertExec := fmt.Sprintf(`INSERT INTO public."USER" (username, email, password) VALUES ('%s', '%s','%s')`, user.Username, user.Email, user.Password)
		fmt.Println(insertExec)
		_, err := db.Exec(insertExec)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func userInvalid(user *models.User) bool {
	return util.IsEmpty(user.Username) ||
		len(user.Username) < 5 ||
		util.IsEmpty(user.Email) ||
		util.IsEmailInvalid(user.Email) ||
		util.IsEmpty(user.Password) ||
		len(user.Password) < 8

}
