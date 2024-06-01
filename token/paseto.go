package token

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/ed25519"
)

type PasetoPayload struct {
	Email              string    `json:"email"`
	IssueDateTime      time.Time `json:"issue_datetime"`
	ExpirationDateTime time.Time `json:"expiration_datetime"`
}

func (pp *PasetoPayload) isValid() bool {
	return time.Now().After(pp.ExpirationDateTime)
}

type PasetoMaker interface {
	CreateToken(email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*PasetoPayload, error)
}

type Paseto struct {
	paseto     *paseto.V2
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

func (p *Paseto) CreateToken(email string, duration time.Duration) (string, error) {
	pp := PasetoPayload{
		Email:              email,
		IssueDateTime:      time.Now(),
		ExpirationDateTime: time.Now().Add(duration),
	}

	return p.paseto.Sign(p.privateKey, pp, nil)
}

func (p *Paseto) VerifyToken(token string) (*PasetoPayload, error) {
	pp := &PasetoPayload{}
	if err := p.paseto.Verify(token, p.publicKey, pp, nil); err != nil {
		return nil, err
	}

	if pp.isValid() {
		return nil, fmt.Errorf("paseto token is expired")
	}

	return pp, nil
}

func NewPasetoMaker() (PasetoMaker, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	pasetoMaker := &Paseto{
		paseto:     paseto.NewV2(),
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	return pasetoMaker, nil
}
