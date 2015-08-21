package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-zoo/bone"
	"github.com/unrolled/render"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// setting logger
	log.WithFields(log.Fields{
		"url_query": r.URL.Query(),
		"url_path":  r.URL.Path,
	}).Info("Getting home view")

	ren := render.New(render.Options{
		Layout: "layout",
	})

	w.Header().Set("Content-Type", "application/json")

	// get stubo uri
	uri := "http://localhost:8001"

	s := &Stubo{&http.Client{}, "localhost", "8001", "http", uri}
	s.getScenariosDetail()
	// response := []byte("Hello!")
	// w.Write(response)
	ren.HTML(w, http.StatusOK, "example", nil)

}

func scenarioDetailedHandler(w http.ResponseWriter, r *http.Request) {

	id := bone.GetValue(r, "id")
	scenario := bone.GetValue(r, "scenario")
	w.Write([]byte("id:" + string(id) + ", scenario: " + scenario))
}
