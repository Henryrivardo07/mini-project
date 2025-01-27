package middlewares

import (
    "log"
    "net/http"
)

// Logger middleware
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Request:", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
