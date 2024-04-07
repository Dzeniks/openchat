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
	"strings"
	"time"
)

func InitChat(r *gin.RouterGroup) {
	chatGroup := r.Group("/chat")
	chatGroup.Use(authorization.AuthRequired())
	{
		chatGroup.POST("/SentPrompt", SentPrompt)
		chatGroup.POST("/SentPromptTest", SentPromptTest)
		chatGroup.POST("/CreateChat", CreateChat)
		chatGroup.POST("/GetChatIdsOfUser", GetChatIdsOfUser)
		chatGroup.POST("/GetChat", GetChat)
	}
}

type GetChatRequest struct {
	ChatId string `json:"chat_id"`
}

type SentPromptRequest struct {
	Prompt string `json:"prompt"`
	ChatID string `json:"chat_id"`
}

type AIRequest struct {
	Input struct {
		Prompts []databaseService.Message `json:"prompts"`
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
		log.Print("ChatID in GOTO: ", chatID)
	}
}

func SentPrompt(r *gin.Context) {
	// Get the request body
	var payloadUser SentPromptRequest
	err := json.NewDecoder(r.Request.Body).Decode(&payloadUser)
	if err != nil {
		r.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	log.Print("Prompt: ", payloadUser.Prompt)
	log.Print("ChatID: ", payloadUser.ChatID)

	accessTokenString := r.GetHeader("Authorization")
	claims := ginUtils.GetClaimsFromToken(r, accessTokenString)

	database := ginUtils.GetDatabase(r)
	var message = databaseService.Message{
		SenderID: claims.UserID,
		Content:  payloadUser.Prompt,
		SentAt:   time.Now(),
	}

	chatID := payloadUser.ChatID
	chat, err := databaseService.GetChatByID(chatID, database)
	if err != nil {
		log.Print(err)
		log.Print("ChatID: ", chatID)
		r.JSON(500, gin.H{"error": err})
		return
	}

	var aiRequest AIRequest = AIRequest{
		Input: struct {
			Prompts []databaseService.Message `json:"prompts"`
		}{
			Prompts: append(chat.Messages, message),
		},
	}

	aiURL := dotEnv.DotEnv.AiUrl
	reqBodyBytes, err := json.Marshal(aiRequest)
	if err != nil {
		log.Print(err)
		r.JSON(500, gin.H{"error": "Error marshaling request body"})
		return
	}

	// Create a new request with POST method, URL, and request body
	req, err := http.NewRequest("POST", aiURL, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		log.Print(err)
		r.JSON(500, gin.H{"error": "Error creating request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", dotEnv.DotEnv.AiAuthKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		r.JSON(500, gin.H{"error": "Error sending request"})
		return
	}

	if resp.StatusCode != 200 {
		log.Print(resp.Status)
		log.Print(resp.StatusCode)

		r.JSON(resp.StatusCode, gin.H{"error": resp.Status})
		return
	}

	var response AIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Print(err)
		r.JSON(500, gin.H{"error": "Error decoding response from AI"})
		return
	}

	var AImessage = databaseService.Message{
		SenderID: "AI",
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

func GetChatIdsOfUser(r *gin.Context) {
	accessTokenString := r.GetHeader("Authorization")
	claims := ginUtils.GetClaimsFromToken(r, accessTokenString)
	database := ginUtils.GetDatabase(r)
	// Get All chat
	chats, err := databaseService.GetAllChats(claims.UserID, database)
	if err != nil {
		log.Print(err)
		r.JSON(500, gin.H{"error": "Error getting chats"})
		return
	}

	// Extract chatID from chat
	var chatIDs []string
	for _, chat := range *chats {
		if len(chat.Messages) == 0 {
			continue
		}
		chatIDs = append(chatIDs, chat.ChatID)
	}

	var chatIDsDate []time.Time
	for _, chat := range *chats {
		if len(chat.Messages) == 0 {
			continue
		}
		chatIDsDate = append(chatIDsDate, chat.Messages[0].SentAt)
	}

	r.JSON(200, gin.H{
		"chat_ids": chatIDs,
		"dates":    chatIDsDate,
	})
}

func GetChat(r *gin.Context) {
	//accessTokenString := r.GetHeader("Authorization")
	//claims := ginUtils.GetClaimsFromToken(r, accessTokenString)
	//userID := claims.UserID
	database := ginUtils.GetDatabase(r)
	var payload GetChatRequest
	err := json.NewDecoder(r.Request.Body).Decode(&payload)
	if err != nil {
		r.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// split ChatID by _
	chatID := strings.Split(payload.ChatId, "_")
	if len(chatID) != 2 {
		r.JSON(400, gin.H{"error": "Invalid ChatID"})
		return
	}

	// Temporary check if user is in chat
	//if chatID[0] != userID {
	//	r.JSON(400, gin.H{"error": "Invalid ChatID"})
	//	return
	//}

	chat, err := databaseService.GetChatByID(payload.ChatId, database)
	if err != nil {
		log.Print(err)
		//r.Error(err)
		r.JSON(404, gin.H{"error": err})
		return
	}
	r.JSON(200, gin.H{
		"chat": chat,
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
		log.Print(err)
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
		log.Print(err)
		r.JSON(500, gin.H{"error": "Error creating chat"})
		return
	}
	log.Print("ChatID: ", *chatID)
	r.JSON(200, gin.H{"chat_id": *chatID})
}
