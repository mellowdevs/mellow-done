package homepage

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/util"
)

func GetHomepage(c *gin.Context) {
	activeUserID := c.Keys["userId"].(string)
	if activeUserID == "" {
		location := url.URL{Path: "api/v1/auth/login"}
		c.Redirect(http.StatusUnauthorized, location.RequestURI())
		return
	} else {
		c.JSON(200, util.ResponseMessage{
			Success: true,
			Message: "Welcome",
		})
	}
}
