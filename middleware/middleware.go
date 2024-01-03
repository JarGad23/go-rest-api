package middleware

import (
	"net/http"

	"github.com/JarGad23/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenicate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
