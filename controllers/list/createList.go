package list

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/models"
	"github.com/mellowdevs/mellow-done/util"
)

func CreateList(c *gin.Context) {
	var db = config.GetDBInstance()
	newList := models.List{}
	newList.CreatedAt = time.Now().Unix()
	newList.UpdatedAt = time.Now().Unix()
	newList.UserID = c.Keys["userId"].(string)

	if bindErr := c.ShouldBindJSON(&newList); bindErr != nil {
		c.JSON(400, util.ResponseMessage{
			Success: false,
			Message: bindErr.Error(),
		})
		return
	}

	strFields := []string{"title", "user_id"}
	intFields := []string{"created_at", "updated_at"}
	strValues := []string{newList.Title, newList.UserID}
	intValues := []int64{newList.CreatedAt, newList.UpdatedAt}

	insertStr := util.GenerateInsertIntoTableStringWithIntValues("list", strFields, intFields, strValues, intValues)
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
		Message: util.LIST_CREATE_SUCCESS,
	})

}
