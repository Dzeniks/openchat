package ginUtils

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"server-go/lib/databaseService"
	"server-go/lib/jwtService"
)

func GetClaimsFromToken(r *gin.Context, accessTokenString string) *jwtService.Claims {
	accessToken, err := jwtService.ParseToken(accessTokenString)
	if err != nil {
		r.JSON(500, gin.H{"error": "Cannot parse token"})
	}
	claims, err := jwtService.ExtractClaims(accessToken)
	if err != nil {
		r.JSON(500, gin.H{"error": "Cannot extract claims"})
	}
	return claims
}

func GetDatabase(r *gin.Context) *mongo.Database {
	client, err := databaseService.GetClient()
	if err != nil {
		r.JSON(500, gin.H{"error": "Cant connect to client"})
	}
	return databaseService.GetDatabase(client)
}
