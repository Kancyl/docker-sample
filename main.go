package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"

	_ "github.com/newrelic/go-agent/v3/integrations/nrmssql"
)

const (
	DB_USER     = "socfin-id\\kdl"
	DB_PASSWORD = "5GH#*TeZ"
	DB_NAME     = "bangunbandar_20221114"
	DB_PORT     = "1433"
	SERVER      = "192.168.7.48"
)

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;database=%s;password=%s;port=%s", SERVER, DB_USER, DB_NAME, DB_PASSWORD, DB_PORT)
	db, err := sqlx.Open("mssql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Simple health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("OK"))
	})

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":9090", nil))
}
