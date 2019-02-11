package prometheus

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/metrics", prometheus.Handler().ServeHTTP)

	return r
}
