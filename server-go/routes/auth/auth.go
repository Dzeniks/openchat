package auth

import (
	"context"
	"github.com/labstack/gommon/log"
	"server-go/lib/authorization"
	"server-go/lib/databaseService"
	"server-go/lib/jwtService"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitAuth RegisterRoutes registers routes for the authorization package.
func InitAuth(r *gin.RouterGroup) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/", auth)
		authGroup.POST("/login", login)
		authGroup.POST("/register", register)
		authGroup.POST("/refresh", refresh)
	}
}

type EmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func register(c *gin.Context) {
	log.Print("register")
	// Define a struct to hold the request body

	// Bind the request body to the RegisterRequest struct
	var req EmailRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// Check if data are valid
	if !authorization.IsValidEmail(req.Email) {
		c.JSON(400, gin.H{"error": "Invalid email"})
		return
	}
	if !authorization.IsValidPassword(req.Password) {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}
	client, err := databaseService.GetClient()
	if err != nil {
		log.Error("Database error Client error")
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		}
	}(client, ctx)
	database := databaseService.GetDatabase(client)
	user, err := databaseService.GetUserByEmail(&req.Email, database)
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	if user != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}
	hashedPassword, err := authorization.HashPassword(req.Password)
	if err != nil {
		log.Error("Password hashing error")
		c.JSON(500, gin.H{"error": "Password hashing error"})
		return
	}
	ID := uuid.New()
	user = &databaseService.User{
		UserID:           ID.String(),
		Username:         req.Email,
		Email:            req.Email,
		Password:         hashedPassword,
		RegistrationDate: time.Now(),
		Active:           true}
	err = databaseService.CreateUser(user, database)
	if err != nil {
		log.Error("Database error")
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	// c.JSON(200, gin.H{"message": "User created"})
	// login after registration
	token, err := jwtService.GenerateJWT(user)
	if err != nil {
		log.Error("JWT generation error")
		c.JSON(500, gin.H{"error": "JWT generation error"})
		return
	}
	refreshToken, err := jwtService.GenerateRefreshToken(user, database)
	if err != nil {
		log.Error("Refresh token generation error")
		c.JSON(500, gin.H{"error": "Refresh token generation error"})
		return
	}
	c.JSON(200, gin.H{"message": "Login successful", "accessToken": token, "refreshToken": refreshToken})
}

func login(c *gin.Context) {
	log.Print("login")
	var req EmailRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if data are valid
	if !authorization.IsValidEmail(req.Email) {
		c.JSON(400, gin.H{"error": "Invalid email"})
		return
	}
	if !authorization.IsValidPassword(req.Password) {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	client, err := databaseService.GetClient()
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		}
	}(client, ctx)
	database := databaseService.GetDatabase(client)

	user, err := databaseService.GetUserByEmail(&req.Email, database)
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	if user == nil {
		c.JSON(400, gin.H{"error": "User does not exist"})
		return
	}
	if !user.Active {
		c.JSON(400, gin.H{"error": "User is not active"})
		return
	}
	if !authorization.ComparePasswords(user.Password, req.Password) {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}
	token, err := jwtService.GenerateJWT(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "JWT generation error"})
		return
	}
	refreshToken, err := jwtService.GenerateRefreshToken(user, database)
	if err != nil {
		c.JSON(500, gin.H{"error": "Refresh token generation error"})
		return
	}
	c.JSON(200, gin.H{"message": "Login successful", "accessToken": token, "refreshToken": refreshToken})
}

func refresh(c *gin.Context) {
	var authToken = c.Request.Header.Get("RefreshToken")

	if authToken == "" {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// IsValid?
	jwtToken, err := jwtService.ParseToken(authToken)
	if err != nil {
		c.JSON(500, gin.H{"error": "JWT parsing error"})
		return
	}
	if !jwtToken.Valid {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	// Check claims
	claims, err := jwtService.ExtractClaims(jwtToken)
	if err != nil {
		c.JSON(500, gin.H{"error": "JWT parsing error"})
		return
	}
	if claims.Type != "refresh" || claims.UserID == "" {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	// Connect to the database
	client, err := databaseService.GetClient()
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	ctx := context.Background()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		}
	}(client, ctx)
	database := databaseService.GetDatabase(client)

	// Check if the user exists and is active
	user, err := databaseService.GetUserByID(&claims.UserID, database)
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	if user == nil {
		c.JSON(400, gin.H{"error": "User does not exist"})
		return
	}
	if !user.Active {
		c.JSON(400, gin.H{"error": "User is not active"})
		return
	}

	accessToken, err := jwtService.GenerateJWT(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "JWT generation error"})
		return
	}
	refreshToken, err := jwtService.GenerateRefreshToken(user, database)
	if err != nil {
		c.JSON(500, gin.H{"error": "Refresh token generation error"})
		return
	}
	c.JSON(200, gin.H{"message": "Refresh successful", "accessToken": accessToken, "refreshToken": refreshToken})
}

func auth(c *gin.Context) {
	var authToken = c.Request.Header.Get("Authorization")

	if authToken == "" {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// IsValid?
	jwtToken, err := jwtService.ParseToken(authToken)
	if err != nil {
		c.JSON(500, gin.H{"error": "JWT parsing error"})
		return
	}
	if !jwtToken.Valid {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	// Check claims
	claims, err := jwtService.ExtractClaims(jwtToken)
	if err != nil {
		c.JSON(500, gin.H{"error": "JWT parsing error"})
		return
	}
	if claims.Type != "access" || claims.UserID == "" {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	// Connect to the database
	client, err := databaseService.GetClient()
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	ctx := context.Background()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		}
	}(client, ctx)
	database := databaseService.GetDatabase(client)

	// Check if the user exists and is active
	user, err := databaseService.GetUserByID(&claims.UserID, database)
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	if user == nil {
		c.JSON(400, gin.H{"error": "User does not exist"})
		return
	}
	if !user.Active {
		c.JSON(400, gin.H{"error": "User is not active"})
		return
	}
	c.JSON(200, gin.H{"message": "Auth successful"})
}
