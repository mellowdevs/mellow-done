package util

import (
	"net/http"
	"net/mail"
)

type ResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var (
	EMPTY_FIELD_MSG    = "cannot be empty."
	SHORT_FIELD_MSG    = "is too short."
	INVALID_FIELD_MSG  = "is invalid."
	REG_SUCCESS        = "Successfully registered."
	PAYLOAD_ERR        = "Payload is required"
	USER_NOT_FOUND     = "User not found with this email."
	INCORRECT_PASSWORD = "Password is not correct."
)

func IsEmpty(str string) bool {
	return len(str) == 0
}
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func IsEmailInvalid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func BadRequest(rw http.ResponseWriter, message string) {
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write([]byte(message))
}
