package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"main/handlers"
	"main/tools"
	"net/http"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type application struct {
	auth struct {
		username string
		password string
	}
}

func main() {
	app := new(application)
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
	mux.HandleFunc("/get_sum", app.basicAuth(handlers.GetSumOfArticlesHandler))
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", mux))
}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
