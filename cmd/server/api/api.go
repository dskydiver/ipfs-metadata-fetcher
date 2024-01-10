package api

import (
	"log"
	"net/http"

	"github.com/dskydiver/ipfs-metadata-fetcher/cmd/server/controllers"
	"github.com/dskydiver/ipfs-metadata-fetcher/cmd/server/routers"
	"github.com/dskydiver/ipfs-metadata-fetcher/pkg/database"
	"github.com/gorilla/mux"
)

func init() {
	log.Println("Initializing api server...")
}

func RunServer() {
	log.Println("Running api server...")

	client := database.NewClient()
	defer client.Close()

	controller := controllers.NewTokenController(client)

	r := mux.NewRouter()

	routers.TokenRoutes(r, controller)

	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
