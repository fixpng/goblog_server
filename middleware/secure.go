package middleware

import github.com/gin-gonic/gin

func Secure() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("X-Frame-Options", "SAMEORIGIN")
        c.Header("X-Content-Type-Options", "nosniff")
        c.Header("X-XSS-Protection", "1; mode=block")
        c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        c.Next()
    }
}
