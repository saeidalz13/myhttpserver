package routes

import (
	"myserver/api/handlers"
	"myserver/api/middleware"
	"net/http"
)

func Setup(mux *http.ServeMux, apiH *handlers.ApiHandler) {
	mux.HandleFunc("POST /signup", apiH.AuthHandler.HandlePostSignUp)
	mux.HandleFunc("POST /login", apiH.AuthHandler.HandlePostLogin)

	mux.HandleFunc("POST /store", middleware.Log(apiH.StoreHandler.HandlePostItem))
	mux.HandleFunc("GET /retrieve/{key}", middleware.Log(apiH.StoreHandler.HandleGetItem))
	mux.HandleFunc("DELETE /delete-item/{key}", middleware.Log(apiH.StoreHandler.HandleDeleteItem))
}
