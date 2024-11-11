package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-espanol/backend/utils"
	"github.com/leetcode-espanol/backend/routes/user"

	"github.com/leetcode-espanol/backend/models"
	"gorm.io/gorm"
)

func migrate_db(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Account{},
		&models.Session{},
		&models.User{},
		&models.VerificationToken{},
		&models.Problem{},
		&models.Solution{},
	)
	if err != nil {
		return err
	}
	return nil
}
func main() {

	db, err := utils.InitDB ()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	migrate_db(db)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	r := gin.Default()
	rg := r.Group("/users")

	user_routes.AddUserRoutes(rg)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(":6000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
