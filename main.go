package main

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"main/handlers"
	"main/tools"
	"net/http"
)

func main() {
	tools.SetupLogging()

	db := tools.DbConn()
	handlers.SetDB(db)
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexPageHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", mux))
}
