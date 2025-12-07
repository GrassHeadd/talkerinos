package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"talkerinos/internal/database"
	"talkerinos/internal/handler"
	"talkerinos/internal/router"
)

func main() {
	_ = godotenv.Load(".env")

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	h := handler.New(database.New(conn))
	r := router.New(h)

	fmt.Println("Server starting on port:", portStr)
	r.Run(":" + portStr)
}
