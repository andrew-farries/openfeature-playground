package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/open-feature/go-sdk/openfeature"
)

type contextKey string

var userMap = map[string]map[string]interface{}{
	"alice": {
		"Identifier": "alice",
		"Email":      "alice@acme.com",
		"Country":    "UK",
	},
	"bob": {
		"Identifier": "bob",
		"Email":      "bob@foo.com",
		"Country":    "Germany",
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

		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		evalContext := openfeature.NewEvaluationContext(token, userMap[token])
		r = r.WithContext(context.WithValue(r.Context(), contextKey("evalContext"), evalContext))

		next.ServeHTTP(w, r)
	})
}
