package task

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func DeleteTask(c *gin.Context) {
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
	_, err := models.GetTaskById(taskReq.Id, activeUserID)
	if err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.TASK_NOT_FOUND,
		})
		return
	}

	sqlStatement := `DELETE FROM task WHERE id = $1;`
	res, deleteErr := db.Exec(sqlStatement, taskReq.Id)
	if deleteErr != nil {
		if deleteErr == sql.ErrNoRows {
			c.JSON(400, util.ResponseMessage{
				Success: false,
				Message: deleteErr.Error(),
			})
			return
		}
		c.JSON(500, util.ResponseMessage{
			Success: false,
			Message: deleteErr.Error(),
		})
		return
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(200, util.ResponseMessage{
			Success: false,
			Message: util.LIST_NOT_FOUND,
		})
		return
	}
	c.JSON(200, util.ResponseMessage{
		Success: true,
		Message: util.TASK_DELETE_SUCCESS,
	})

}
