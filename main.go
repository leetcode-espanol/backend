package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func init_db() *gorm.DB {
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
	return db
}

func main() {

	db := init_db()
	db.AutoMigrate(&Product{})
	e := echo.New()

	e.GET("/vivo", func(c echo.Context) error {
		//return c.String(http.StatusOK, "está vivooooooooo")
		val, err := json.Marshal(map[string]string{"message": "está vivooooooooo"}) 
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err))  
		}
		return c.JSON(http.StatusOK, val) 

	})
	if err := e.Start(":6000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
