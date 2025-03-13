package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

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
