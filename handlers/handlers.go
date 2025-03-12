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

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	switch r.Method {
	case "GET":
		log.Infof("Get request %v", r.Body)
		customRsp := &core.CustomHttpReponse{Form: core.Form{Title: "Title", Body: "Ла ла ла"}}
		t.Execute(w, customRsp)
	case "POST":
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		log.Info(r.PostForm)
		customRsp := &core.CustomHttpReponse{Form: core.Form{Title: r.FormValue("title"), Body: r.FormValue("body")}}
		log.Infof("Post request form data %v", r.PostForm)
		t.Execute(w, customRsp)
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
