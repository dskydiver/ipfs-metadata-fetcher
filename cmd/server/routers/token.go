package routers

import (
	"net/http"

	"github.com/dskydiver/ipfs-metadata-fetcher/cmd/server/controllers"
	"github.com/gorilla/mux"
)

func TokenRoutes(r *mux.Router, controller *controllers.TokenController) {
	r.HandleFunc("/tokens", controller.ListTokens).Methods(http.MethodGet)
	r.HandleFunc("/tokens/{cid}", controller.GetToken).Methods(http.MethodGet)
}
