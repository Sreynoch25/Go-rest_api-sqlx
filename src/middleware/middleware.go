package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// UserContext represents the structure for storing user information
type UserContext struct {
	Email        string
	Exp          time.Time
	LoginSession string
	RoleId       float64
	UserID       int
	UserName     string
	UserAgent    string
	Ip           string
}

// JwtMiddleware is a middleware to extract and store user context from JWT
func JwtMiddleware() fiber.Handler {
	// Load the secret key from environment variables
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		panic("JWT_SECRET is not set in the environment variables")
	}

	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		// Split the Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization format",
			})
		}

		tokenString := parts[1]

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token is signed with HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(secretKey), nil
		})

		// Check if it invalid
		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Extract token to claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		// Extract user information for user context
		// handle error one by one to avoid nil pointer
		// claim email
		email, ok := claims["email"].(string)
		if !ok {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "email is missing",
			})
		}
		// claim login session
		login_session, ok := claims["login_session"].(string)
		if !ok {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "login is missing",
			})
		}

		// claim role id
		role_id, ok := claims["role_id"].(float64)
		if !ok {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "role id is missing",
			})
		}

		// claim user id
		user_id, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "user id is missing",
			})
		}

		//claim username
		user_name, ok := claims["user_name"].(string)
		if !ok {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "user name is missing",
			})
		}

		// claim exp
		exp, ok := claims["exp"].(float64)
		if !ok {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "user name is missing",
			})
		}

		// Store user context in locals of go
		// it will clear out when go return the value to client
		c.Locals("UserContext", UserContext{
			Email:        email,
			Exp:          time.Unix(int64(exp), 0),
			LoginSession: login_session,
			RoleId:       role_id,
			UserID:       int(user_id),
			UserName:     user_name,
			// user agent is what client use for request into our backend 
			// example : firefox , insomnia ........
			UserAgent:    string(c.Context().UserAgent()),
			// it detect the ip that request into backend
			Ip:           c.Context().LocalIP().String(),
		})

		return c.Next()
	}
}