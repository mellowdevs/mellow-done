package list

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func UpdateList(c *gin.Context) {
	var db = config.GetDBInstance()
	activeUserID := c.Keys["userId"].(string)
	listReq := models.List{}
	listReq.UpdatedAt = time.Now().Unix()
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

	sqlStatement := `UPDATE list SET title = $2, updated_at = $3 WHERE id = $1;`
	res, updateErr := db.Exec(sqlStatement, listReq.Id, listReq.Title, listReq.UpdatedAt)
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
			Message: util.LIST_NOT_FOUND,
		})
		return
	}
	c.JSON(200, util.ResponseMessage{
		Success: true,
		Message: util.LIST_UPDATE_SUCCESS,
	})

}
