package chat

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"server-go/lib/authorization"
	"server-go/lib/databaseService"
	"server-go/lib/dotEnv"
	"server-go/lib/ginUtils"
	"time"
)

func InitChat(r *gin.RouterGroup) {
	chatGroup := r.Group("/chat")
	chatGroup.Use(authorization.AuthRequired())
	{
		chatGroup.POST("/SentPrompt", SentPrompt)
		chatGroup.POST("/SentPromptTest", SentPromptTest)
		//chatGroup.GET("/Chat")
		chatGroup.POST("/CreateChat", CreateChat)
	}
}

type GetChatRequest struct {
	ChatId string
}

type SentPromptRequest struct {
	Prompt string `json:"prompt"`
	ChatID string `json:"chat_id"`
}

type AIRequest struct {
	Input struct {
		Prompts []string `json:"prompts"`
	} `json:"input"`
}

type AIResponse struct {
	Output string `json:"output"`
}

func addMessageToChatParallel(message databaseService.Message,
	chatID string, database *mongo.Database) {
	err := databaseService.AddMessageToChat(message, chatID, database)
	if err != nil {
		log.Print(err)
	}
}

func SentPrompt(r *gin.Context) {

	var payloadUser SentPromptRequest
	log.Println(r.Request.Body)
	err := json.NewDecoder(r.Request.Body).Decode(&payloadUser)
	if err != nil {
		log.Println("Error decoding request body:", err)
		return
	}
	accessTokenString := r.GetHeader("Authorization")
	claims := ginUtils.GetClaimsFromToken(r, accessTokenString)

	database := ginUtils.GetDatabase(r)

	var message = databaseService.Message{
		//ChatID
		SenderID: claims.UserID,
		Content:  payloadUser.Prompt,
		SentAt:   time.Now(),
	}

	//Get chat
	chatID := payloadUser.ChatID
	chat, err := databaseService.GetChatByID(chatID, database)
	if err != nil {
		r.JSON(500, gin.H{"error": err.Error()})
		return
	}
	chat.Messages = append(chat.Messages, message)
	var chatMessages []string
	for _, message := range chat.Messages {
		chatMessages = append(chatMessages, message.Content)
	}

	payloadAI := AIRequest{
		Input: struct {
			Prompts []string `json:"prompts"`
		}{
			Prompts: chatMessages,
		},
	}

	//Request to AI_URL
	aiURL := dotEnv.DotEnv.AiUrl
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
	log.Print(resp.StatusCode)
	if resp.StatusCode != 200 {
		r.JSON(resp.StatusCode, gin.H{"error": "Error"})
		return
	}

	var response AIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("Error decoding response body:", err)
		r.JSON(500, gin.H{"error": "Error decoding response from AI"})
		return
	}

	var AImessage = databaseService.Message{
		SenderID: "0",
		Content:  response.Output,
		SentAt:   time.Now(),
	}

	go addMessageToChatParallel(message, payloadUser.ChatID, database)
	go addMessageToChatParallel(AImessage, payloadUser.ChatID, database)

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

func CreateChat(r *gin.Context) {
	accessTokenString := r.GetHeader("Authorization")
	claims := ginUtils.GetClaimsFromToken(r, accessTokenString)

	database := ginUtils.GetDatabase(r)

	// Get All chat
	chats, err := databaseService.GetAllChats(claims.UserID, database)
	if err != nil {
		r.JSON(500, gin.H{"error": "Error getting chats"})
		return
	}

	//If exists chat with no messages return chatID
	for _, chat := range *chats {
		if len(chat.Messages) == 0 {
			log.Println("Chat with no messages", chat.ChatID)
			r.JSON(200, gin.H{"chat_id": chat.ChatID})
			return
		}
	}

	//Create list with element claims.UserID
	var userIds []string
	userIds = append(userIds, claims.UserID)

	chatID, err := databaseService.CreateChat(claims.UserID, userIds, database)
	if err != nil {
		r.JSON(500, gin.H{"error": "Error creating chat"})
		return
	}
	log.Print("chatID: ", *chatID)
	r.JSON(200, gin.H{"chat_id": *chatID})
}
