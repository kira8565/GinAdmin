package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Get(c)
		println("==========")
		println(session.Get("isLogin") == nil)
		println("==========")
		if session.Get("isLogin") == nil && c.Request.URL.Path != "/" {
			c.Redirect(http.StatusMovedPermanently, "/?msg=请登录")
		}
	}
}
