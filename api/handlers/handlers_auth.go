package handlers

import (
	"encoding/json"
	"io"
	"log"
	"myserver/api/db"
	"myserver/api/internal/handlerutils"
	"myserver/api/models"
	"myserver/api/token"
	"net/http"
	"strings"
	"time"
)

type AuthHandler struct {
	Db          *db.Db
	PasetoMaker token.PasetoMaker
}

func NewAuthHandler(db *db.Db, pasetoMaker token.PasetoMaker) *AuthHandler {
	return &AuthHandler{
		Db:          db,
		PasetoMaker: pasetoMaker,
	}
}

func (a *AuthHandler) HandlePostSignUp(w http.ResponseWriter, r *http.Request) {
	if !handlerutils.IsContentTypeJson(r.Header.Get("Content-Type")) {
		http.Error(w, "invalid type of content in post request", http.StatusBadRequest)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, ErrUnableProcessReq, http.StatusServiceUnavailable)
		return
	}

	var user models.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		http.Error(w, ErrInvalidUserJson, http.StatusBadRequest)
		return
	}

	if err := handlerutils.ValidatePassword(user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password and normalize the email
	hashedPassword, err := handlerutils.GenerateHashedPassword(user.Password)
	if err != nil {
		http.Error(w, ErrUnableProcessReq, http.StatusServiceUnavailable)
		return
	}
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		http.Error(w, ErrNoEmail, http.StatusBadRequest)
		return
	}

	// Add user to database
	if err := a.Db.InsertUser(user.Email, hashedPassword); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookieToken, err := a.PasetoMaker.CreateToken(user.Email, time.Hour)
	if err != nil {
		log.Println(err)
		http.Error(w, ErrAuth, http.StatusBadRequest)
		return
	}

	authResp := models.AuthResp{
		Token: cookieToken,
	}
	resp, _ := json.Marshal(authResp)

	// If the browser called us
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "paseto_auth",
	// 	Value:    cookieToken,
	//  Path:     "/"
	// 	Expires:  time.Now().Add(time.Hour),
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	SameSite: http.SameSiteStrictMode,
	// })
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (a *AuthHandler) HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	if !handlerutils.IsContentTypeJson(r.Header.Get("Content-Type")) {
		http.Error(w, "invalid type of content in post request", http.StatusBadRequest)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, ErrUnableProcessReq, http.StatusServiceUnavailable)
		return
	}

	var user models.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		http.Error(w, ErrInvalidUserJson, http.StatusBadRequest)
		return
	}

	hashedPassword, err := a.Db.SelectUser(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := handlerutils.ComparePasswords(hashedPassword, user.Password); err != nil {
		log.Println(err)
		http.Error(w, ErrWrongPassword, http.StatusBadRequest)
		return
	}

	cookieToken, err := a.PasetoMaker.CreateToken(user.Email, time.Hour)
	if err != nil {
		log.Println(err)
		http.Error(w, ErrAuth, http.StatusBadRequest)
		return
	}

	authResp := models.AuthResp{
		Token: cookieToken,
	}
	resp, _ := json.Marshal(authResp)

	// If the browser called us
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "paseto_auth",
	// 	Value:    cookieToken,
	//  Path:     "/"
	// 	Expires:  time.Now().Add(time.Hour),
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	SameSite: http.SameSiteStrictMode,
	// })
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
