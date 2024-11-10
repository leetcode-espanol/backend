package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/leetcode-espanol/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func init_db() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	if host == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
		host = os.Getenv("DB_HOST")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		db_name = os.Getenv("DB_NAME")
		port = os.Getenv("DB_PORT")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, db_name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}


	    err = db.AutoMigrate(
        &models.Account{},
        &models.Session{},
        &models.User{},
        &models.VerificationToken {},
    )
    if err != nil {
        return nil, err
    }


	return db, nil
}

func main() {

	_, err := init_db()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	e := echo.New()

	e.GET("/vivo", func(c echo.Context) error {

		json.NewEncoder(c.Response().Writer).Encode(map[string]string{"message":"está vivooooooooo"})
		return nil

	})
	if err := e.Start(":6000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
