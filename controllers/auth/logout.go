package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mellowdevs/mellow-done/util"
)

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("token", "", 2000, "/", "", true, true)
	c.JSON(200, util.ResponseMessage{
		Success: true,
		Message: util.LOGOUT_SUCCESS,
	})
}
