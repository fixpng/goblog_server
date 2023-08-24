package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"gvb_server/global"
)

// TlsHandler 中间件开启https，还没捣鼓好
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     global.Config.System.Addr(),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
