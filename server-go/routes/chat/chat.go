package chat

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"server-go/lib/dotEnv"
	"github.com/gin-gonic/gin"
)

type ChatCompletetionRequest struct {
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

func InitChat(r *gin.RouterGroup) {
	chatGroup := r.Group("/chat")
	{
		chatGroup.POST("/ChatCompletetion", ChatCompletetion)
		chatGroup.POST("/ChatCompletetionTest", ChatCompletetionTest)
	}
}

func ChatCompletetion(r *gin.Context) {

	var payload ChatCompletetionRequest
	err := json.NewDecoder(r.Request.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding request body:", err)
		return
	}

	payloadAI := AIRequest{
		Input: struct {
			Prompt string `json:"prompt"`
		}{
			Prompt: payload.Prompt,
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
		log.Println("Error making request to AI_URL:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("AI_URL returned non-200 status code:", resp.StatusCode)
		return
	}
	var chatCompletetionResponse AIResponse
	err = json.NewDecoder(resp.Body).Decode(&chatCompletetionResponse)
	if err != nil {
		log.Println("Error decoding response body:", err)
		return
	}
	r.JSON(200, gin.H{
		"message": "OK",
		"output":  chatCompletetionResponse.Output,
	})

}

func ChatCompletetionTest(r *gin.Context) {
	dummy := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl vitae aliquam lacinia, nunc nisl luctus nunc, vitae aliquam nisl nunc eu nisl."

	r.JSON(200, gin.H{
		"message": "OK",
		"output":  dummy,
	})
}
