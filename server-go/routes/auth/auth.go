package auth

import (
	"context"
	"log"
	"regexp"
	"server-go/lib/databaseService"
	"server-go/lib/jwtService"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterRoutes registers routes for the auth package.
func InitAuth(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", login)
		auth.POST("/register", Register)
		auth.POST("/refresh", Refresh)
	}
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func isValidEmail(email string) bool {
	// Regular expression for a simple email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	// Use the MatchString method to check if the email matches the pattern
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	// Check minimum length
	if len(password) < 8 {
		return false
	}

	// Check for at least one uppercase letter
	upperCaseRegex := regexp.MustCompile(`[A-Z]`)
	if !upperCaseRegex.MatchString(password) {
		return false
	}

	// Check for at least one digit
	digitRegex := regexp.MustCompile(`[0-9]`)
	if !digitRegex.MatchString(password) {
		return false
	}

	// All checks passed
	return true
}

func Register(c *gin.Context) {
	// Define a struct to hold the request body

	// Bind the request body to the RegisterRequest struct
	var req AuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// Check if data are valid
	if !isValidEmail(req.Email) {
		c.JSON(400, gin.H{"error": "Invalid email"})
		return
	}
	if !isValidPassword(req.Password) {
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
	if user != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}
	hashedPassword, err := jwtService.HashPassword(req.Password)
	if err != nil {
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
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	c.JSON(200, gin.H{"message": "User created"})
}

func login(c *gin.Context) {
	var req AuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if data are valid
	if !isValidEmail(req.Email) {
		c.JSON(400, gin.H{"error": "Invalid email"})
		return
	}
	if !isValidPassword(req.Password) {
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
	log.Print(user.UserID)
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
	if !jwtService.ComparePasswords(user.Password, req.Password) {
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
	c.JSON(200, gin.H{"message": "Login successful", "token": token, "refreshToken": refreshToken})
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func Refresh(c *gin.Context) {
	// Bind the request body to the RefreshRequest struct
	var req RefreshRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if req.RefreshToken == "" {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// IsValid?
	jwtToken, err := jwtService.ParseToken(req.RefreshToken)
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
		log.Print(claims.Type)
		log.Print(claims.UserID)
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
	c.JSON(200, gin.H{"message": "Refresh successful", "token": accessToken, "refreshToken": refreshToken})
}
