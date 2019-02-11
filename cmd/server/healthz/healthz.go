package healthz

import (
	"net/http"

	"github.com/gorilla/mux"
)

type healthzHandler struct{}

func (h *healthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func newHealthz() *healthzHandler {
	return &healthzHandler{}
}

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", newHealthz().ServeHTTP)

	return r
}
