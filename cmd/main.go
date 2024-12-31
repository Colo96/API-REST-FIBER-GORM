package main

import (
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
		log.Fatalf("Error resolving .env file path: %v", err)
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file")
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
		log.Fatalf("Could not connect to the database: %v", err)
	}

	err = models.MigrateUsers(db)

	if err != nil {
		log.Fatal("could not migrate the database", err)
	}

	app := fiber.New()
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
