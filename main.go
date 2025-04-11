package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/mini-api/config"
	"github.com/sidiqPratomo/mini-api/internal/handler"
	"github.com/sidiqPratomo/mini-api/internal/repository"
	"github.com/sidiqPratomo/mini-api/internal/usecase"
)

func main() {
	// Initialize DB
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}
	defer db.Close()

	// Create users table (auto migrate)
	createUserTable(db)

	// Initialize Repository & Usecase
	userRepo := repository.NewUserRepo(db)
	userUC := usecase.NewUserUsecase(userRepo)

	// Setup router
	r := gin.Default()

	// Register routes
	handler.NewUserHandler(r, userUC)

	// Start server
	fmt.Println("🚀 Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}

func createUserTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal("❌ Failed to create users table:", err)
	}

	fmt.Println("✅ User table checked/created successfully")
}
