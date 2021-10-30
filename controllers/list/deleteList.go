package list

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func DeleteList(c *gin.Context) {
	var db = config.GetDBInstance()
	activeUserID := c.Keys["userId"].(string)
	listReq := models.List{}
	if bindErr := c.ShouldBindJSON(&listReq); bindErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: bindErr.Error(),
		})
		return
	}
	_, err := models.GetListById(listReq.Id, activeUserID)
	if err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.LIST_NOT_FOUND,
		})
		return
	}
	if deleteTasksErr := models.DeleteTasksByListID(listReq.Id, activeUserID); err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: deleteTasksErr.Error(),
		})
		return
	}

	sqlStatement := `DELETE FROM list WHERE id = $1;`
	res, deleteErr := db.Exec(sqlStatement, listReq.Id)
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
		Message: util.LIST_DELETE_SUCCESS,
	})

}
