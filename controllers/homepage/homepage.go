package homepage

import (
	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/util"
)

func GetHomepage(c *gin.Context) {


	c.JSON(200, util.ResponseMessage{
		Success: true,
		Message: "Welcome!",
	})
	return
}
