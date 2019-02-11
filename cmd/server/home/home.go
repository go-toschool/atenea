package home

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	return r
}
