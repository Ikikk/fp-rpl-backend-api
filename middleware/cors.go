package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Controll-Allow-Origin", "*")
		c.Header("Access-Controll-Allow-Credentials", "true")
		c.Header("Access-Controll-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-")
		c.Header("Access-Controll-Allow-Methods", "POST,HEAD,OPTIONS,GET,PUT,DELETE")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
