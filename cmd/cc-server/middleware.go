package main

import (
	"context"
	"net/http"
	"strings"

	configcat "github.com/configcat/go-sdk/v9"
)

type contextKey string

var userMap = map[string]*configcat.UserData{
	"alice": {
		Identifier: "alice",
		Email:      "alice@acme.com",
		Country:    "UK",
	},
	"bob": {
		Identifier: "bob",
		Email:      "bob@foo.com",
		Country:    "Germany",
	},
}

func FeatureFlagMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		token, ok := strings.CutPrefix(authHeader, "Bearer ")
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, ok := userMap[token]
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), contextKey("userData"), user))

		next.ServeHTTP(w, r)
	})
}
