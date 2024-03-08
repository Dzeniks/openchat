package chat

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server-go/lib/auth"
	"server-go/lib/dotEnv"
)

func InitChat(r *gin.RouterGroup) {
	chatGroup := r.Group("/chat")
	chatGroup.Use(auth.AuthRequired())
	{
		chatGroup.POST("/SentPrompt", SentPrompt)
		chatGroup.POST("/SentPromptTest", SentPromptTest)
		//chatGroup.GET("/Chat")
	}
}

type GetChatRequest struct {
	ChatId string
}

func GetChat(r *gin.Context) {
}

type SentPromptRequest struct {
	Prompt string `json:"prompt"`
}

type AIRequest struct {
	Input struct {
		Prompt string `json:"prompt"`
	} `json:"input"`
}

type AIResponse struct {
	Output string `json:"output"`
}

func SentPrompt(r *gin.Context) {
	var payloadUser SentPromptRequest
	err := json.NewDecoder(r.Request.Body).Decode(&payloadUser)
	if err != nil {
		log.Println("Error decoding request body:", err)
		return
	}

	// TODO: Save in parallel to database

	payloadAI := AIRequest{
		Input: struct {
			Prompt string `json:"prompt"`
		}{
			Prompt: payloadUser.Prompt,
		},
	}

	// Request to AI_URL
	aiURL := dotEnv.DotEnv.AI_URL
	reqBodyBytes, err := json.Marshal(payloadAI)
	if err != nil {
		log.Println("Error marshaling request body:", err)
		return
	}
	resp, err := http.Post(aiURL, "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		r.JSON(500, gin.H{"error": "Error making request to AI_URL"})
		return
	}
	if resp.StatusCode != 200 {
		r.JSON(resp.StatusCode, gin.H{"message": "Error"})
		return
	}

	var response AIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("Error decoding response body:", err)
		r.JSON(500, gin.H{"error": "Error decoding response from AI"})
		return
	}

	// TODO: Save in parallel to database

	r.JSON(200, gin.H{
		"message": "OK",
		"output":  response.Output,
	})

}

func SentPromptTest(r *gin.Context) {
	dummy := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl vitae aliquam lacinia, nunc nisl luctus nunc, vitae aliquam nisl nunc eu nisl."

	r.JSON(200, gin.H{
		"message": "OK",
		"output":  dummy,
	})
}
