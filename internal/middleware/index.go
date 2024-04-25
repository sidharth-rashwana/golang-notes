package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckMiddleware(c *gin.Context) {
	headers := c.GetHeader("Authorization")
	fmt.Println(headers)

	if headers == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Headers not provided",
		})
		return
	}
}
