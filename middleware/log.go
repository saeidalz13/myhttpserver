package middleware

import (
	"log"
	"net/http"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("got request with path: %s", r.URL.Path)

        // Forward the request to the handler
        next.ServeHTTP(w, r)
    }
}
