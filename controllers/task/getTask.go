package task

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func GetTaskById(c *gin.Context) {
	taskReq := models.Task{}
	activeUserID := c.Keys["userId"].(string)
	if bindErr := c.ShouldBindJSON(&taskReq); bindErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: bindErr.Error(),
		})
		return
	}
	task, err := models.GetTaskById(taskReq.Id, activeUserID)
	if err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.TASK_NOT_FOUND,
		})
		return
	}

	taskArr := []models.Task{task}

	c.JSON(200, util.ResponseTaskItem{
		Success: true,
		Message: taskArr,
	})
}

func GetTasksByListId(c *gin.Context) {
	var db = config.GetDBInstance()
	taskReq := models.Task{}
	activeUserID := c.Keys["userId"].(string)
	if bindErr := c.ShouldBindJSON(&taskReq); bindErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: bindErr.Error(),
		})
		return
	}
	_, err := models.GetListById(taskReq.ListID, activeUserID)
	if err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.LIST_NOT_FOUND,
		})
		return
	}

	whereMap := make(map[string]string)
	whereMap["user_id"] = activeUserID
	whereMap["list_id"] = taskReq.ListID
	selectStr := util.GenerateSelectFromTableString("task", whereMap)
	rows, err := db.Query(selectStr)
	if err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.LIST_NOT_FOUND,
		})
		return
	}
	defer rows.Close()
	taskArr := []models.Task{}
	for rows.Next() {
		task := models.Task{}
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.ListID, &task.UserID, &task.CreatedAt, &task.UpdatedAt); err != nil {
			c.JSON(400, util.ResponseMessage{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		taskArr = append(taskArr, task)
	}

	c.JSON(200, util.ResponseTaskItem{
		Success: true,
		Message: taskArr,
	})
}

func GetAllTasks(c *gin.Context) {
	var db = config.GetDBInstance()
	activeUserID := c.Keys["userId"].(string)
	whereMap := make(map[string]string)
	whereMap["user_id"] = activeUserID
	selectStr := util.GenerateSelectFromTableString("task", whereMap)
	rows, err := db.Query(selectStr)
	if err != nil {

		fmt.Println("rows")
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.TASK_NOT_FOUND,
		})
		return
	}
	defer rows.Close()
	taskArr := []models.Task{}
	for rows.Next() {
		task := models.Task{}
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.ListID, &task.UserID, &task.CreatedAt, &task.UpdatedAt); err != nil {
			c.JSON(400, util.ResponseMessage{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		taskArr = append(taskArr, task)
	}

	c.JSON(200, util.ResponseTaskItem{
		Success: true,
		Message: taskArr,
	})
}
