package auth

import "net/http"

func Logout(rw http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(rw, &http.Cookie{
		Name:  "token",
		Value: "",
	})
}
