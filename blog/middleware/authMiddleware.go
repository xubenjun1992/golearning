package middleware

import (
	"blog/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(401, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}
		jwtToken, err := utils.ParseToken(token)

		if err != nil {
			ctx.JSON(401, gin.H{"error": "Invalid token: " + err.Error()})
			ctx.Abort()
			return
		}

		if !jwtToken.Valid {
			ctx.JSON(401, gin.H{"error": "Token is invalid"})
			ctx.Abort()
			return
		}
		if claims, ok := jwtToken.Claims.(jwt.MapClaims); !ok {
			ctx.JSON(401, gin.H{"error": "Invalid token claims"})
			ctx.Abort()
			return
		} else {
			ctx.Set("userId", claims["userId"])
			log.Printf("User ID from token: %v", claims["userId"])
			ctx.Set("userName", claims["userName"])
			ctx.Next()
		}
	}
}
