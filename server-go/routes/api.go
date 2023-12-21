package routes

import (
	"github.com/gin-gonic/gin"

	"server-go/routes/auth"
)

func InitApiRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	auth.InitAuth(api)

}
