package session

import (
	"net/http"

	"github.com/go-toschool/palermo/auth"
	"github.com/go-toschool/sicily/cmd/server/cors"
	"github.com/go-toschool/syracuse/citizens"
	"github.com/gorilla/mux"
)

func Routes(ctx *Context) *mux.Router {
	r := mux.NewRouter()

	cs := ctx.Handle(CreateSession)
	ds := ctx.Handle(DeleteSession)

	cors := cors.NewCors()
	r.HandleFunc("/session", cors.Check(cs)).Methods("POST")
	r.HandleFunc("/session", cors.Check(ds)).Methods("DELETE")

	return r
}

type AuthCors struct {
	UserService    citizens.CitizenshipClient
	SessionService auth.AuthServiceClient
}

// CheckCors ...
func (ac *AuthCors) CheckCors(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
