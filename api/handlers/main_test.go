package handlers

import (
	"myserver/api/db"
	"myserver/api/token"
	"net/http"
	"os"
	"testing"
	"time"
)

var HttpClientTest http.Client
var StoreHandlerTest *StoreHandler
var AuthHandlerTest *AuthHandler

func TestMain(m *testing.M) {
	dbMemory := db.NewDb()
	pasetoMaker, err := token.NewPasetoMaker()
	if err != nil {
		panic(err)
	}

    // Handlers
	sh := NewStoreHandler(dbMemory)
	StoreHandlerTest = sh

	au := NewAuthHandler(dbMemory, pasetoMaker)
	AuthHandlerTest = au

    // Client test struct
	c := http.Client{
		Timeout: time.Second * 5,
	}
	HttpClientTest = c
	os.Exit(m.Run())
}
