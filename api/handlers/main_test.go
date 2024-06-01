package handlers

import (
	"myserver/api/db"
	"myserver/api/token"
	"os"
	"testing"
)

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

	os.Exit(m.Run())
}
