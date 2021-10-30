package task

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func UpdateTask(c *gin.Context) {
	var db = config.GetDBInstance()
	activeUserID := c.Keys["userId"].(string)
	taskReq := models.Task{}
	taskReq.UpdatedAt = time.Now().Unix()
	if bindErr := c.ShouldBindJSON(&taskReq); bindErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: bindErr.Error(),
		})
		return
	}
	if _, err := models.GetTaskById(taskReq.Id, activeUserID); err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.TASK_NOT_FOUND,
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

	sqlStatement := `UPDATE task SET title = $2, description = $3, list_id = $4, status = $5,  updated_at = $6 WHERE id = $1;`
	res, updateErr := db.Exec(sqlStatement, taskReq.Id, taskReq.Title, taskReq.Description, taskReq.ListID, taskReq.Status, taskReq.UpdatedAt)
	if updateErr != nil {
		if updateErr == sql.ErrNoRows {
			c.JSON(400, util.ResponseMessage{
				Success: false,
				Message: updateErr.Error(),
			})
			return
		}
		c.JSON(500, util.ResponseMessage{
			Success: false,
			Message: updateErr.Error(),
		})
		return
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(200, util.ResponseMessage{
			Success: false,
			Message: util.TASK_NOT_FOUND,
		})
		return
	}
	c.JSON(200, util.ResponseMessage{
		Success: true,
		Message: util.TASK_UPDATE_SUCCESS,
	})

}
