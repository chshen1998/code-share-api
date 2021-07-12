package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DB *sql.DB
)

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port, _ := strconv.Atoi(os.Getenv("DBPORT"))
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", os.Getenv("DBHOST"), port, os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))
	DB, _ = sql.Open("postgres", connStr)
}
