package handlers

import (
	"database/sql"
	"main/core"
	"net/http"
	"text/template"

	log "github.com/sirupsen/logrus"
)

var db *sql.DB

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	switch r.Method {
	case "GET":
		log.Infof("Get request %v", r.Body)
		customRsp := core.Form{Title: "Title", Body: "Message"}
		t.Execute(w, customRsp)
	case "POST":
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		log.Info(r.PostForm)
		customRsp := core.Form{Title: r.FormValue("title"), Body: r.FormValue("body")}
		sqlStmt := `INSERT INTO user_requests (title, message) VALUES ($1, $2)`
		_, err = db.Exec(sqlStmt, customRsp.Title, customRsp.Body)
		if err != nil {
			panic(err)
		}
		log.Infof("Post request form data %v", r.PostForm)
		t.Execute(w, customRsp)
	}
}

func PathDbToHandlers(dbConn *sql.DB) {
	db = dbConn
}
