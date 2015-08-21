package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// setting logger
	log.WithFields(log.Fields{
		"url_query": r.URL.Query(),
		"url_path":  r.URL.Path,
	}).Info("Getting home view")

	w.Header().Set("Content-Type", "application/json")
	response := []byte("Hello!")
	w.Write(response)

}
