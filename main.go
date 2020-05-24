package main

import (
	"log"
	"net/http"

	"github.com/theAndrewCline/it-resource-manager/postgres"
	"github.com/theAndrewCline/it-resource-manager/web"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@postres/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":4000", h)
}
