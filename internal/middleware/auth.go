package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/laudryfadian/starter-golang/internal/domain"
	"github.com/laudryfadian/starter-golang/internal/helpers"
)

type AuthMiddleware struct {
	userUsecase domain.UserRepository
}

func NewAuthMiddleware(userUsecase domain.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		userUsecase: userUsecase,
	}
}

func (a *AuthMiddleware) AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.ErrorResponse(c, "Authorization header is required", http.StatusUnauthorized)
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			helpers.ErrorResponse(c, "Invalid authorization format", http.StatusUnauthorized)
			c.Abort()
			return
		}

		tokenString := bearerToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			helpers.ErrorResponse(c, "Invalid token", http.StatusUnauthorized)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helpers.ErrorResponse(c, "Invalid token claims", http.StatusUnauthorized)
			c.Abort()
			return
		}

		_, err = a.userUsecase.GetByID(c.Request.Context(), claims["userId"].(string))
		if err != nil {
			helpers.ErrorResponse(c, "User not found", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("userId", claims["userId"])
		c.Set("email", claims["email"])
		c.Next()
	}
}
