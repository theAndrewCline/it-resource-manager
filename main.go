package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/theAndrewCline/it-resource-manager/postgres"
	"github.com/theAndrewCline/it-resource-manager/web"

	// postgres drivers
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "postgres"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	store, err := postgres.NewStore(connStr)
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":4000", h)
}
