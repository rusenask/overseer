package main

import (
	"flag"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
	"github.com/meatballhat/negroni-logrus"
)

func main() {
	// looking for option args when starting App
	// like ./lgc -port=":3000" would start on port 3000
	var port = flag.String("port", ":3000", "Server port")
	flag.Parse() // parse the flag

	log.WithFields(log.Fields{
		"port": port,
	}).Info("Overseer is starting")

	mux := bone.New()
	mux.Get("/", http.HandlerFunc(homeHandler))
	mux.Get("/stubos/:id/scenarios/:scenario", http.HandlerFunc(scenarioDetailedHandler))
	n := negroni.Classic()
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(mux)
	n.Run(*port)
}
