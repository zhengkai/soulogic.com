package main

import (
	"net/http"
	"time"
)

func server() (ok bool) {

	mux := http.NewServeMux()
	mux.HandleFunc(`/gateway`, handleGateway)

	s := &http.Server{
		Addr:         httpListenAddr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// server.SetKeepAlivesEnabled(false)

	j(`start web server`, httpListenAddr)

	err := s.ListenAndServe()
	if err != nil {
		jw(`web server fail:`, err)
		return false
	}

	return true
}
