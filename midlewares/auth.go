package midlewares

import (
	"fmt"
	"net/http"

	"golang_basic_gin_mei/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		fmt.Println("token: ", tokenString)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Request Need Access Token",
			})

			c.Abort()
			return
		}

		// validate token
		_, _, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})

			c.Abort()
			return
		}

		c.Next()

	}
}
