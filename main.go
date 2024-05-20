package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/tools"
	"net/http"
)

func main() {
	tools.SetupLogging()

	db := tools.GetDbConnection()
	handlers.PathDbToHandlers(db)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Errorf("Error closing db connection: %v", err)
		}
	}(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexPageHandler)
	mux.HandleFunc("/get_sum", handlers.GetSumOfArticlesHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", mux))
}
