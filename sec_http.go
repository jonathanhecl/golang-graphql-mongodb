package main

import (
	"context"
	"net/http"
)

// Header Authorization
func headerAuthorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		ctx = context.WithValue(r.Context(), "username", "") // default empty
		auth := r.Header.Get("authorization")
		if auth != "" {
			token, _ := verifyToken(auth)
			if token != nil {
				ctx = context.WithValue(r.Context(), "username", token["username"])
			}
		}
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Disable CORS
func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
