package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sidharth-rashwana/notes-app/internal/utils"

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

	token := strings.Split(headers, " ")
	if len(token) < 2 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Token not provided",
		})
		return
	}

	err := utils.TokenCheck(token[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "claims not matched",
		})
		return
	}
	c.Next()
}
