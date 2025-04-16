package handlers

import (
	"encoding/json"
	"main/core"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func DataRepresent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ges := make([]core.Row, 0)
		sqlStmt := `SELECT COUNT(id),  DATE_TRUNC('hour', pub_date) as dt FROM UserRequests GROUP BY dt ORDER BY dt`
		var row core.Row
		rows, errSql := db.Query(sqlStmt)
		if errSql != nil {
			http.Error(w, errSql.Error(), 400)
			return
		}

		for rows.Next() {
			rows.Scan(&row.Count, &row.Dt)
			ges = append(ges, row)
		}
		lables := make([]string, 0)
		var ds core.DataSet
		var ge core.GraphicElement
		var data []int
		for i := 0; i < len(ges); i++ {
			el := &ges[i]
			lables = append(lables, el.Dt.String())
			log.Debugf("each el: %v", el)
			data = append(data, el.Count)
		}
		ds = core.DataSet{Label: "Group by hours", Data: data}
		ge.Datasets = append(ge.Datasets, ds)
		ge.Labels = lables
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		jsonResp, _ := json.Marshal(ge)
		log.Debug(ges)
		w.Write(jsonResp)
	}
}
