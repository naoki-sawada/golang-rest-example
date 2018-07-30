package middleware

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		fmt.Printf("auth: %s \n", auth)

		if auth == "" {
			c.String(401, "Authorization header required")
			c.Abort()
			return
		}

		splitStr := strings.Split(auth, " ")
		if splitStr[0] != "Bearer" {
			c.String(401, "Authorization is not Bearer type")
			c.Abort()
			return
		}

		if splitStr[1] == "" {
			c.String(401, "Bearer token not found")
			c.Abort()
			return
		}

		tokenString := splitStr[1]
		token, e := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("token-token"), nil
		})

		if !token.Valid {
			c.String(401, "Auth token failed: %s", e)
			c.Abort()
			return
		}
	}
}
