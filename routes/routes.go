package routes

import (
	"myserver/handlers"
	"myserver/middleware"
	"net/http"
)

func Setup(mux *http.ServeMux, sh *handlers.ServerHandler) {
	mux.HandleFunc("POST /store", middleware.Log(sh.HandlePostItem))
	mux.HandleFunc("GET /retrieve/{key}", middleware.Log(sh.HandleGetItem))
	mux.HandleFunc("DELETE /delete-item/{key}", middleware.Log(sh.HandleDeleteItem))
}
