package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-toschool/atenea"
	"github.com/go-toschool/palermo/auth"
	"github.com/go-toschool/sicily"
	"github.com/go-toschool/syracuse/citizens"
)

type SessionData struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Token    string `json:"token"`
}

// CreateSession creates a new authentication token.
func CreateSession(ctx *Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This server does not support that HTTP method", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read body", http.StatusInternalServerError)
		return
	}

	var sd SessionData

	if err := json.Unmarshal(body, &sd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctxb := context.Background()
	opts := &citizens.CreateRequest{
		Data: &citizens.Citizen{
			Email:    sd.Email,
			FullName: sd.FullName,
		},
	}
	// create user in syracusa service.
	u, err := ctx.User.Create(ctxb, opts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, iat, err := idWithTime()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sess := &auth.CreateRequest{
		Data: &auth.Session{
			Id:        id,
			UserId:    u.Data.Id,
			Email:     sd.Email,
			Token:     sd.Token,
			CreatedAt: iat.Unix(),
			UpdatedAt: iat.Unix(),
		},
	}
	// create access token in palermo service.
	sessCred, err := ctx.Session.Create(ctxb, sess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &sicily.User{
		ID:       u.Data.Id,
		Email:    u.Data.Email,
		FullName: u.Data.FullName,
		Token:    sessCred.Data.AuthToken,
	}

	w.Header().Set("Accept-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func idWithTime() (string, *time.Time, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", nil, err
	}

	iat := time.Now()
	id := base64.StdEncoding.EncodeToString(b)

	return id, &iat, nil
}

// DeleteSession delete an active session
func DeleteSession(ctx *Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This server does not support that HTTP method", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read body", http.StatusInternalServerError)
		return
	}

	var sd SessionData

	if err := json.Unmarshal(body, &sd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctxb := context.Background()
	opts := &citizens.CreateRequest{
		Data: &citizens.Citizen{
			Email:    sd.Email,
			FullName: sd.FullName,
		},
	}
	// create user in syracusa service.
	u, err := ctx.User.Create(ctxb, opts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, iat, err := idWithTime()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sess := &auth.CreateRequest{
		Data: &auth.Session{
			Id:        id,
			UserId:    u.Data.Id,
			Email:     sd.Email,
			Token:     sd.Token,
			CreatedAt: iat.Unix(),
			UpdatedAt: iat.Unix(),
		},
	}
	// create access token in palermo service.
	sessCred, err := ctx.Session.Create(ctxb, sess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &atenea.User{
		ID:       u.Data.Id,
		Email:    u.Data.Email,
		FullName: u.Data.FullName,
		Token:    sessCred.Data.AuthToken,
	}

	w.Header().Set("Accept-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
