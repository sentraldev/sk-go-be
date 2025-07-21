package middleware

import (
	"context"
	"net/http"

	auth "firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

type FirebaseClient struct {
	authClient *auth.Client
}

func NewFirebaseClient(authClient *auth.Client) *FirebaseClient {
	return &FirebaseClient{authClient: authClient}
}

func (fc *FirebaseClient) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Expect header format: "Bearer <token>"
		token := authHeader[len("Bearer "):]
		verifiedToken, err := fc.authClient.VerifyIDToken(context.Background(), token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			c.Abort()
			return
		}

		// Store user ID in context for use in handlers
		c.Set("uid", verifiedToken.UID)

		c.Next()
	}
}
