package main

import (
	"fmt"
	"net/http"

	configcat "github.com/configcat/go-sdk/v9"
)

type server struct {
	mux      *http.ServeMux
	ccClient *configcat.Client
}

func NewServer(client *configcat.Client) *server {
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
	user := r.Context().Value(contextKey("userData")).(configcat.User)

	isMyFirstFeatureEnabled := s.ccClient.GetBoolValue("isMyFirstFeatureEnabled", false, user)
	if isMyFirstFeatureEnabled {
		fmt.Fprintln(w, "feature is enabled")
	} else {
		fmt.Fprintln(w, "feature is disabled")
	}

	isSomeOtherFeatureEnabled := s.ccClient.GetBoolValue("someOtherFlag", false, user)
	if isSomeOtherFeatureEnabled {
		fmt.Fprintln(w, "local-only feature is enabled")
		return
	} else {
		fmt.Fprintln(w, "local-only feature is disabled")
	}
}
