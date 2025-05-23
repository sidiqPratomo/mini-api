package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("❌ Could not load .env file")
	}
    host := os.Getenv("DB_HOST")
    portStr := os.Getenv("DB_PORT")
    port, err := strconv.Atoi(portStr)
    if err != nil {
		return nil, fmt.Errorf("invalid port number: %v", err)
    }
	
    user := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_DATABASE")
	
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
	user, password, host, port, dbname)
	
	fmt.Println("DSN :", dsn)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to open DB: %v", err)
    }

    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping DB: %v", err)
    }

    fmt.Println("✅ Connected to MySQL database")
    return db, nil
}

func GetDB() *sql.DB {
	return DB
}
