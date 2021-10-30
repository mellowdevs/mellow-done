package task

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func CreateTask(c *gin.Context) {
	var db = config.GetDBInstance()
	newTask := models.Task{}
	newTask.CreatedAt = time.Now().Unix()
	newTask.UpdatedAt = time.Now().Unix()
	newTask.Status = util.FIRST_STATUS
	newTask.UserID = c.Keys["userId"].(string)

	if bindErr := c.ShouldBindJSON(&newTask); bindErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: bindErr.Error(),
		})
		return
	}

	strFields := []string{"title", "description", "status", "list_id", "user_id"}
	intFields := []string{"created_at", "updated_at"}
	strValues := []string{newTask.Title, newTask.Description, newTask.Status, newTask.ListID, newTask.UserID}
	intValues := []int64{newTask.CreatedAt, newTask.UpdatedAt}

	insertStr := util.GenerateInsertIntoTableStringWithIntValues("task", strFields, intFields, strValues, intValues)
	fmt.Println(insertStr)
	_, insertErr := db.Exec(insertStr)
	if insertErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: insertErr.Error(),
		})
		return
	}
	c.JSON(200, util.ResponseMessage{
		Success: true,
		Message: util.TASK_CREATE_SUCCESS,
	})

}
