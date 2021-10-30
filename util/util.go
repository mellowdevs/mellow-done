package util

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/mellowdevs/mellow-done/models"
)

type ResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type ResponseUserEvent struct {
	Success bool                `json:"success"`
	Message models.UserResponse `json:"message"`
}

type ResponseListItem struct {
	Success bool          `json:"success"`
	Message []models.List `json:"message"`
}
type ResponseTaskItem struct {
	Success bool          `json:"success"`
	Message []models.Task `json:"message"`
}

var (
	EMPTY_FIELD_MSG     = "cannot be empty."
	SHORT_FIELD_MSG     = "is too short."
	INVALID_FIELD_MSG   = "is invalid."
	REG_SUCCESS         = "Successfully registered."
	LOGOUT_SUCCESS      = "Successfully logged out."
	PAYLOAD_ERR         = "Payload is required"
	USER_NOT_FOUND      = "User not found with this email."
	INCORRECT_PASSWORD  = "Password is not correct."
	NOT_AUTHED          = "Not authorized"
	NOT_LOGGED          = "Not logged in"
	EMAIL_EXIST         = "This email already exists"
	USERNAME_EXIST      = "This username already exists"
	LIST_CREATE_SUCCESS = "Successfully created the list"
	LIST_UPDATE_SUCCESS = "Successfully updated the list"
	LIST_DELETE_SUCCESS = "Successfully deleted the list"
	LIST_NOT_FOUND      = "List not found"
	TASK_CREATE_SUCCESS = "Successfully created the task"
	TASK_UPDATE_SUCCESS = "Successfully updated the task"
	TASK_DELETE_SUCCESS = "Successfully deleted the task"
	TASK_NOT_FOUND      = "Task not found"
	FIRST_STATUS        = "Undone"
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
	rw.WriteHeader(http.StatusUnauthorized)
	rw.Write([]byte(message))
}

func GenerateSelectFromTableString(table string, whereMap map[string]string) string {
	var fieldsStr string = "*"
	var whereStr string = ""

	for k, v := range whereMap {
		if IsNotEmpty(whereStr) {
			whereStr += " AND "
		}
		var statement string = fmt.Sprintf((k + " = '%s'"), v)
		whereStr += statement
	}
	selectStr := fmt.Sprintf(`SELECT %s FROM public."%s" WHERE %s`, fieldsStr, table, whereStr)
	return selectStr
}

func GenerateInsertIntoTableString(table string, fields []string, values []string) string {
	if len(fields) != len(values) {
		return ""
	}
	fieldsString := GenerateListingString(fields)
	valuesString := ""
	for index, value := range values {
		valuesString += fmt.Sprintf(`'%s'`, value)
		if index != len(values)-1 {
			valuesString += ", "
		}
	}
	insertStr := fmt.Sprintf(`INSERT INTO public."%s" (%s) VALUES (%s)`, table, fieldsString, valuesString)

	return insertStr
}

func GenerateInsertIntoTableStringWithIntValues(table string, strFields []string, intFields []string, strValues []string, intValues []int64) string {
	if (len(strFields) != len(strValues)) || (len(intFields) != len(intValues)) {
		return ""
	}

	fieldsString := GenerateListingString(strFields)
	if len(intFields) > 0 {
		fieldsString += ", "
		fieldsString += GenerateListingString(intFields)
	}
	valuesString := ""
	for index, value := range strValues {
		valuesString += fmt.Sprintf(`'%s'`, value)
		if index != len(strValues)-1 {
			valuesString += ", "
		}
	}
	if len(intValues) > 0 {
		valuesString += ", "
		for index, value := range intValues {
			valuesString += fmt.Sprintf(`%v`, value)
			if index != len(intValues)-1 {
				valuesString += ", "
			}
		}
	}
	insertStr := fmt.Sprintf(`INSERT INTO public."%s" (%s) VALUES (%s)`, table, fieldsString, valuesString)

	return insertStr
}

func GenerateListingString(list []string) string {
	var result string = ""
	for index, item := range list {
		result += item
		if index != len(list)-1 {
			result += ", "
		}
	}
	return result
}
