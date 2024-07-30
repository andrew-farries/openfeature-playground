package main

import (
	"fmt"
	"net/http"

	"github.com/open-feature/go-sdk/openfeature"
)

type server struct {
	mux      *http.ServeMux
	ccClient *openfeature.Client
}

func NewServer(client *openfeature.Client) *server {
	s := &server{
		mux:      http.NewServeMux(),
		ccClient: client,
	}

	s.mux.HandleFunc("GET /hello", s.handleHello)
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := FeatureFlagMiddleware(s.mux)

	mux.ServeHTTP(w, r)
}

func (s *server) handleHello(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(contextKey("evalContext")).(openfeature.EvaluationContext)

	isMyFirstFeatureEnabled, _ := s.ccClient.BooleanValue(r.Context(), "isMyFirstFeatureEnabled", false, user)

	if isMyFirstFeatureEnabled {
		fmt.Fprintln(w, "feature is enabled")
	} else {
		fmt.Fprintln(w, "feature is disabled")
	}

	isSomeOtherFeatureEnabled, _ := s.ccClient.BooleanValue(r.Context(), "someOtherFlag", false, user)
	if isSomeOtherFeatureEnabled {
		fmt.Fprintln(w, "local-only feature is enabled")
		return
	} else {
		fmt.Fprintln(w, "local-only feature is disabled")
	}
}
