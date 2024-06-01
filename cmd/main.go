package main

import (
	"log"
	"myserver/api/db"
	"myserver/api/handlers"
	"myserver/api/routes"
	"myserver/api/token"
	"net/http"
)

func main() {
	// starting a server multiplexer
	mux := http.NewServeMux()

	// set up in-memory for the handlers
	dbMemory := db.NewDb()

	// Paseto maker
	pasetoMaker, err := token.NewPasetoMaker()
	if err != nil {
		panic(err)
	}

	apiHandler := &handlers.ApiHandler{
		StoreHandler: handlers.NewStoreHandler(dbMemory),
		AuthHandler:  handlers.NewAuthHandler(dbMemory, pasetoMaker),
	}

	// setting up the routes
	routes.Setup(mux, apiHandler)

	// port will be in .env file
	log.Println("listening to port 2024")
	log.Fatalln(http.ListenAndServe(":2024", mux))
}
