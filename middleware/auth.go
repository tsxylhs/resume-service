package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// CorsHandler 跨域请求
func CorsHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("origin")
		//Logger.Debug("request from ", zap.String("origin", origin))

		if strings.ToUpper(method) == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Mx-ReqToken,X-Requested-With,Agent,access-token")
			c.Writer.Header().Set("Access-Control-Max-Age", "1728000")
			c.AbortWithStatus(204)
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
		return
	}
}
