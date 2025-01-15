package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/webdav"
)

func logger(r *http.Request, err error) {
	if err != nil {
		fmt.Printf("Request: %s, Method: %s, Error: %v\n", r.URL.Path, r.Method, err)
	} else {
		fmt.Printf("Request: %s, Method: %s\n", r.URL.Path, r.Method)
	}
}

func digestAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		println(authHeader)
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", "Digest realm=\"Webdav\"")
			http.Error(w, "Unauthorized : Authorization header is empty", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := &webdav.Handler{
		FileSystem: webdav.Dir(`/tmp/test`),
		LockSystem: webdav.NewMemLS(),
		Logger:     logger,
	}

	print("Start WebDav server")
	http.ListenAndServe("localhost:8080", digestAuthMiddleware(handler))
}
