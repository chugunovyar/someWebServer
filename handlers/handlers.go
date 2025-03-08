package handlers

import (
	"database/sql"
	"encoding/json"
	"main/core"
	"net/http"
	"text/template"

	log "github.com/sirupsen/logrus"
)

var db *sql.DB

const format = "2006-01-02 15:04:05"

func loadPage() *core.Page {
	return &core.Page{Title: "Мощный заголовок", Body: "ла-ла-ла"}
}

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	switch r.Method {
	case "GET":
		p := loadPage()
		log.Infof("Get request %v", r.Body)
		t.Execute(w, p)
	case "POST":
		p := loadPage()
		log.Infof("Post request form data %v", r.FormValue("body"))
		t.Execute(w, p)
	}
}

func GetSumOfArticlesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		sqlStmt := `SELECT COUNT(*) FROM ARTICLES`
		var count int
		errSql := db.QueryRow(sqlStmt).Scan(&count)
		if errSql != nil {
			http.Error(w, errSql.Error(), 400)
			return
		}
		log.Debugf("count of articles: %d", count)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		jsonResp, _ := json.Marshal(map[string]int{"count": count})
		w.Write(jsonResp)
	}

}

func PathDbToHandlers(dbConn *sql.DB) {
	db = dbConn
}
