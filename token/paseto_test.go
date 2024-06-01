package token

import (
	"fmt"
	"testing"
	"time"
)

func PasetoTest(t *testing.T) {
	pasetoMaker, err := NewPasetoMaker()
	if err != nil {
		t.Fatal(err)
	}

	email := "some_email@gmail.com"
	token, err := pasetoMaker.CreateToken(email, time.Hour)
    if err != nil {
        t.Fatal(err)
    }
    fmt.Printf("generated token: %s", token)

    pp, err := pasetoMaker.VerifyToken(token)
    if err != nil {
        t.Fatal(err)
    }
    fmt.Printf("Decrypted Payload: %+v", pp)

    if pp == nil {
        t.Fatal("failed to authenticate the user")
    }

    if pp.Email != email {
        t.Fatal("emails don't match after decryption")
    }
}
