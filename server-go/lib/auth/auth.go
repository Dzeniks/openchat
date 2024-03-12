package auth

import (
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
	"server-go/lib/jwtService"
)

type AuthRequest struct {
	AccessToken string `json:"access_token"`
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if authenticated
		authToken := c.GetHeader("Authorization")
		log.Print("authToken: ", authToken)
		if authToken == "" {
			c.Abort()
		}
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

		c.Next()
	}
}

func IsValidEmail(email string) bool {
	// Regular expression for a simple email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	// Use the MatchString method to check if the email matches the pattern
	return re.MatchString(email)
}

func IsValidPassword(password string) bool {
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
