package chat

import (
	"log"
	"net/http"
	"server-go/lib/dotEnv"
	"github.com/gin-gonic/gin"
)

func InitChat(r *gin.RouterGroup) {
	chatGroup := r.Group("/chat")
	{
		chatGroup.POST("/ChatCompletetion", ChatCompletetion)
		chatGroup.POST("/ChatCompletetionTest", ChatCompletetionTest)
	}
}

func ChatCompletetion(r *gin.Context) {
	aiURL := dotEnv.DotEnv.AI_URL

	resp, err := http.Get(aiURL)
	if err != nil {
		log.Println("Error making request to AI_URL:", err)
		return
	}
	defer resp.Body.Close()

}

func ChatCompletetionTest(r *gin.Context) {
	dummy := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl vitae aliquam lacinia, nunc nisl luctus nunc, vitae aliquam nisl nunc eu nisl."

	r.JSON(200, gin.H{
		"message": "OK",
		"output":  dummy,
	})
}
