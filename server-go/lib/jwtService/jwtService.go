package jwtService

import (
	"log"
	"server-go/lib/dotEnv"
	"time"

	"server-go/lib/databaseService"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Generate a salt and hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword string, password string) bool {
	// Compare the hashed password with the password provided by the user
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func GenerateJWT(user *databaseService.User) (*string, error) {
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Token expiration time (1 hour in this example)
		"iat":     time.Now().Unix(),                    // Issued At time
		"type":    "access",                             // Token type
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(dotEnv.DotEnv.SecretKey))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}

type Claims struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(dotEnv.DotEnv.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractClaims(token *jwt.Token) (*Claims, error) {
	// Extract the claims from the refresh token
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	tokenType := claims["type"].(string)
	return &Claims{
		UserID: userID,
		Type:   tokenType,
	}, nil
}

func GenerateRefreshToken(user *databaseService.User, database *mongo.Database) (*string, error) {
	// Define the claims for the refresh token
	log.Print(user.UserID)
	claims := jwt.MapClaims{
		"user_id": user.UserID,                                // Subject (user ID)
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(), // Expiration time (e.g., 30 days)
		"iat":     time.Now().Unix(),                          // Issued At time
		"type":    "refresh",                                  // Token type
	}
	// Sign the token with a secret key
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := refreshToken.SignedString([]byte(dotEnv.DotEnv.SecretKey))
	if err != nil {
		return nil, err
	}
	user.RefreshToken = tokenString
	err = databaseService.InsertRefreshToken(user, database)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
