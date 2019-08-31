package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	w.Write([]byte(requestVars["shortenID"]))
}
