package handlers

import (
	"database/sql"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"main/core"
	"main/tools"
	"net/http"
)

var db *sql.DB

const format = "2006-01-02 15:04:05"

type RespBody struct {
	id int
}

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var article core.Article
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sqlStmt := `INSERT INTO ARTICLES(headline, content, pub_date) VALUES($1,$2,$3) RETURNING id`
		log.Debugln(article)
		var id int
		errSql := db.QueryRow(sqlStmt, article.Headline, article.Content, tools.ConvertTimeToTimestamp(article.PubDate.Format(format))).Scan(&id)
		if errSql != nil {
			http.Error(w, errSql.Error(), 400)
			return
		}
		log.Debugf("Write article id: %d", id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		jsonResp, err := json.Marshal(map[string]int{"id": id})
		w.Write(jsonResp)
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
