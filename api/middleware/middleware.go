package middleware

import "myserver/api/token"

type Middleware struct {
	PasetoMaker token.PasetoMaker
}
