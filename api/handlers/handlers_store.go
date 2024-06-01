package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"myserver/api/db"
	"myserver/api/internal/handlerutils"
	"myserver/api/models"
	"net/http"
)


type StoreHandler struct {
	Db *db.Db
}

func NewStoreHandler(db *db.Db) *StoreHandler {
	return &StoreHandler{
		Db: db,
	}
}

func (s *StoreHandler) HandleDeleteItem(w http.ResponseWriter, r *http.Request) {
	key := handlerutils.ExtractKey(r.URL.Path)

	if err := s.Db.DeleteItem(key); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *StoreHandler) HandleGetItem(w http.ResponseWriter, r *http.Request) {
	// I tried PathValue but seems like it's not compatible with
	// httptest library. I manually extract the key here
	key := handlerutils.ExtractKey(r.URL.Path)

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

func (s *StoreHandler) HandlePostItem(w http.ResponseWriter, r *http.Request) {
	if !handlerutils.IsContentTypeJson(r.Header.Get("Content-Type")) {
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
