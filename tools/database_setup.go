package tools

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func DbConn() *sql.DB {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var dbErr error
	db, dbErr := sql.Open("postgres", psqlConn)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	return db
}
