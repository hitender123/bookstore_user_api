package user_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hitender123/bookstore_user_api/utils/common"
	"github.com/joho/godotenv"
)

const (
	USER        = "user"
	PASS        = "pass"
	HOST        = "127.0.0.1:3306"
	DB          = "db"
	MAX_RETRIES = 10
	RETRY_DELAY = 2 * time.Second
)

var (
	Client   *sql.DB
	username = getEnv("DB_USER", USER)
	userpass = getEnv("DB_PASSWORD", PASS)
	host     = getEnv("DB_HOST", HOST)
	schema   = getEnv("DB_NAME", DB)
)

func getEnv(key, defaultValue string) string {
	//-- load env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found, using environment variables: %v", err)
	}

	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func init() {
	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		userpass,
		host,
		schema,
	)
	fmt.Println("dataSource", dataSource)

	// Retry logic with exponential backoff
	for attempt := 1; attempt <= MAX_RETRIES; attempt++ {
		Client, err = sql.Open("mysql", dataSource)
		if err != nil {
			log.Printf("Failed to open database (attempt %d/%d): %v", attempt, MAX_RETRIES, err)
			time.Sleep(RETRY_DELAY)
			continue
		}

		err = Client.Ping()
		if err == nil {
			log.Println("Database connected successfully")
			return
		}

		log.Printf("Failed to ping database (attempt %d/%d): %v", attempt, MAX_RETRIES, err)
		common.SafeClose(Client)
		time.Sleep(RETRY_DELAY)
	}
	panic(fmt.Sprintf("Failed to connect to database after %d attempts: %v", MAX_RETRIES, err))
}
