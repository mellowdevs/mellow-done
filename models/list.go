package models

import (
	"fmt"

	"github.com/mellowdevs/mellow-done/config"
)

type List struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	UserID    string `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func GetListById(id string, user_id string) (List, error) {
	list := List{}
	var db = config.GetDBInstance()
	selectStr := fmt.Sprintf(`SELECT * FROM PUBLIC."list" WHERE id = %s AND user_id = %s`, id, user_id)
	fmt.Println(selectStr)
	if err := db.QueryRow(selectStr).Scan(&list.Id, &list.Title, &list.UserID, &list.CreatedAt, &list.UpdatedAt); err != nil {
		return list, err
	}
	return list, nil
}
