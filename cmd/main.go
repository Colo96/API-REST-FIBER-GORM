package main

import (
	"api-rest-fiber-gorm/src/controllers"
	"api-rest-fiber-gorm/src/database"
	"api-rest-fiber-gorm/src/models"
	"api-rest-fiber-gorm/src/routes"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	envPath, err := filepath.Abs("../.env")
	if err != nil {
		log.Printf("Error resolving .env file path: %v", err)
		return
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Println("Error loading .env file, continuing without it")
	}

	requiredVars := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Printf("Missing required environment variable: %s", v)
			return
		}
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	db, err := database.Connect(config)
	if err != nil {
		log.Printf("Could not connect to the database: %v", err)
		return
	}

	controllers.SetUpDatabase(db)

	if err := models.MigrateUsers(db); err != nil {
		log.Printf("Could not migrate the database: %v", err)
		return
	}

	app := fiber.New()
	routes.Setup(app)

	if err := app.Listen(":3000"); err != nil {
		log.Printf("Failed to start the server: %v", err)
	}
}
