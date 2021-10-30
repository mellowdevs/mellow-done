package list

import (
	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func GetList(c *gin.Context) {
	listReq := models.List{}
	activeUserID := c.Keys["userId"].(string)
	if bindErr := c.ShouldBindJSON(&listReq); bindErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: bindErr.Error(),
		})
		return
	}

	list, err := models.GetListById(listReq.Id, activeUserID)
	if err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.LIST_NOT_FOUND,
		})
		return
	}

	listArr := []models.List{list}

	c.JSON(200, util.ResponseListItem{
		Success: true,
		Message: listArr,
	})

}

func GetAllLists(c *gin.Context) {
	var db = config.GetDBInstance()
	activeUserID := c.Keys["userId"].(string)
	whereMap := make(map[string]string)
	whereMap["user_id"] = activeUserID
	selectStr := util.GenerateSelectFromTableString("list", whereMap)

	rows, err := db.Query(selectStr)
	if err != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: util.LIST_NOT_FOUND,
		})
		return
	}
	defer rows.Close()

	listArr := []models.List{}
	for rows.Next() {
		list := models.List{}
		if err := rows.Scan(&list.Id, &list.Title, &list.UserID, &list.CreatedAt, &list.UpdatedAt); err != nil {
			c.JSON(400, util.ResponseMessage{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		listArr = append(listArr, list)
	}

	c.JSON(200, util.ResponseListItem{
		Success: true,
		Message: listArr,
	})
}
