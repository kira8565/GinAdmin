package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Get(c)
		if session.Get("isLogin") == nil &&
			(c.Request.URL.Path != "/" || c.Request.URL.Path != "/checklogin") {
			c.Redirect(http.StatusPermanentRedirect, "?msg=请登录")
		}
	}
}
