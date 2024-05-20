package tools

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func GetDbConnection() *sql.DB {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("DB_HOST", "localhost"),
		GetEnvAsInt("POSTGRES_PORT", 5432),
		GetEnv("POSTGRES_USER", "postgres"),
		GetEnv("POSTGRES_PASSWORD", "postgres"),
		GetEnv("POSTGRES_NAME", "postgres"))
	var dbErr error
	db, dbErr := sql.Open("postgres", psqlConn)
	if dbErr != nil {
		log.Fatalf("Error connecting to database: %s", dbErr)
	}
	return db
}
