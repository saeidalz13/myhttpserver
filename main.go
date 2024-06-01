package main

import (
	"log"
	"myserver/db"
	"myserver/handlers"
	"myserver/routes"
	"net/http"
)

func main() {
    // starting a server multiplexer
	mux := http.NewServeMux()

    // set up in-memory for the handlers
	db := db.NewDb()
	serverHandler := handlers.NewServerHandler(db)

    // setting up the routes
	routes.Setup(mux, serverHandler)

    // port will be in .env file
    log.Println("listening to port 2024")
	log.Fatalln(http.ListenAndServe(":2024", mux))
}
