package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"myserver/db"
	"myserver/internal/urlutils"
	"myserver/models"
	"net/http"
)

const (
	ErrUnableProcessReq = "sorry! unable to process the request at the moment"
	ErrNoKeyProvidedUrl = "no key was provided in url path"
)

type ServerHandler struct {
	Db *db.Db
}

func NewServerHandler(db *db.Db) *ServerHandler {
	return &ServerHandler{
		Db: db,
	}
}

func (s *ServerHandler) HandleDeleteItem(w http.ResponseWriter, r *http.Request) {
	key, err := urlutils.ExtractKey(r.URL.Path)
	if err != nil {
        log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.Db.DeleteItem(key); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *ServerHandler) HandleGetItem(w http.ResponseWriter, r *http.Request) {
	// I tried PathValue but seems like it's not compatible with
	// httptest library. I manually extract the key here
	key, err := urlutils.ExtractKey(r.URL.Path)
	if err != nil {
        log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, err := s.Db.SelectItem(key)
	if err != nil {
		http.Error(w, fmt.Errorf("%w key: %s", err, key).Error(), http.StatusNotFound)
		return
	}

	respRaw := map[string]string{
		"value": value,
	}
	// TODO: take care of this error
	resp, _ := json.Marshal(respRaw)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s *ServerHandler) HandlePostItem(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "invalid type of content in post request", http.StatusBadRequest)
		return
	}

	req, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, ErrUnableProcessReq, http.StatusServiceUnavailable)
		return
	}

	var incomingReq models.Item
	if err := json.Unmarshal(req, &incomingReq); err != nil {
		http.Error(w, ErrUnableProcessReq, http.StatusBadRequest)
		return
	}

	keyValue, err := s.Db.InsertItem(&incomingReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: handle if json Marshal was unsuccessful, inform the client
	resp, err := json.Marshal(keyValue)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
