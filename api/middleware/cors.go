package middleware

import "net/http"


func Cors(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")  // allow all the origins
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")  // allow all these methods
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // allow these headers
        w.Header().Set("Access-Control-Max-Age", "600")  // Cache the result of the preflight req for 10 mins

        // If it was a preflight request
        if r.Method == "OPTIONS" {
            return
        }

        next.ServeHTTP(w, r)
    }
}