package chat

import (
	"log"
	"net/http"
	"server-go/lib/dotEnv"
	"time"

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
	// Get the AI_URL from the environment variables
	aiURL := dotEnv.DotEnv.AI_URL

	// Make a GET request to the AI_URL
	resp, err := http.Get(aiURL)
	if err != nil {
		log.Println("Error making request to AI_URL:", err)
		return
	}
	defer resp.Body.Close()

	// Handle the response
	// ...

	// Return the response to the client

	// Save the response to the database

}

func ChatCompletetionTest(r *gin.Context) {
	dummy := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl vitae aliquam lacinia, nunc nisl luctus nunc, vitae aliquam nisl nunc eu nisl."

	// Wait for 2 seconds
	time.Sleep(2 * time.Second)

	r.JSON(200, gin.H{
		"message": "OK",
		"output":  dummy,
	})
}
