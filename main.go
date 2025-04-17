package main

import (
	"database/sql"
	"main/handlers"
	"main/tools"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
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
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/api/get_data", handlers.DataRepresent)
	mux.HandleFunc("/api/get_sum", handlers.GetSumOfArticlesHandler)
	mux.HandleFunc("/", handlers.IndexPageHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", mux))
}
