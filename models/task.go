package models

import (
	"fmt"

	"github.com/mellowdevs/mellow-done/config"
)

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ListID      string `json:"list_id"`
	UserID      string `json:"user_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func GetTaskById(id string, user_id string) (Task, error) {
	task := Task{}
	var db = config.GetDBInstance()
	selectStr := fmt.Sprintf(`SELECT * FROM PUBLIC."task" WHERE id = %s AND user_id = %s`, id, user_id)
	if err := db.QueryRow(selectStr).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.ListID, &task.UserID, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return task, err
	}
	return task, nil
}

func DeleteTasksByListID(list_id, user_id string) error {
	var db = config.GetDBInstance()
	deleteStr := `DELETE FROM public."task" WHERE list_id = $1 AND user_id = $2`
	fmt.Println(deleteStr)
	if _, err := db.Exec(deleteStr, list_id, user_id); err != nil {
		return err
	}
	return nil
}
