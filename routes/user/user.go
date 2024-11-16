package user_routes

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-espanol/backend/models"
	"github.com/leetcode-espanol/backend/utils"
	"gorm.io/gorm"
)

func GetUserBySessionToken(db *gorm.DB, sessionToken string) (*models.User, error) {
	var session *models.Session

	err := db.
		Joins("User").
		Where("\"sessions\".\"sessionToken\" = ?", sessionToken).
		First(&session).Error

	if err != nil {
		return nil, err
	}
	if session.Expires.Before(time.Now()) {
		return nil, errors.New("session expired")
	}

	return &session.User, nil
}

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the session token from the Authorization header
		token := c.GetHeader("Authorization")
		token = strings.Split(token, " ")[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		user, err := GetUserBySessionToken(db, token)
		log.Println(user) 
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session token"})
			c.Abort()
			return
		}

		// Attach the user to the context
		c.Set("user", user)
		c.Next()
	}
}

func AddUserRoutes(rg *gin.RouterGroup) {

	db, err := utils.InitDB()
	rg.Use(AuthMiddleware(db))

	rg.GET("/", func(c *gin.Context) {
		if err != nil {
			c.JSON(http.StatusInternalServerError, "failed to connect database")
			return
		}

		user, found := c.Get("user")
		if found {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusUnauthorized, "user not found")
		}
	})

}
