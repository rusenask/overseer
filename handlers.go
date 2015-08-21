package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-zoo/bone"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// setting logger
	log.WithFields(log.Fields{
		"url_query": r.URL.Query(),
		"url_path":  r.URL.Path,
	}).Info("Getting home view")

	w.Header().Set("Content-Type", "application/json")
	s := &Stubo{&http.Client{}, "localhost", "8001", "http"}
	s.getScenariosDetail()
	response := []byte("Hello!")
	w.Write(response)

}

func scenarioDetailedHandler(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	scenario := bone.GetValue(r, "scenario")
	w.Write([]byte("id:" + string(id) + ", scenario: " + scenario))
}
