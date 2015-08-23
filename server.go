package main

import (
	"flag"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
	"github.com/meatballhat/negroni-logrus"
	"github.com/unrolled/render"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// looking for option args when starting App
	// like ./overeer -port=":3000" would start on port 3000
	var port = flag.String("port", ":3000", "Server port")
	var dbActions = flag.String("db", "", "Database actions - create, migrate, drop")
	flag.Parse() // parse the flag

	// init connection
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Fatal("Failed to open sqlite DB")
	}
	defer db.Close()
	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// flag to do something with the database
	if *dbActions != "" {
		log.WithFields(log.Fields{"action": dbActions}).Info("Database action initiated.")
		d := DBActions{db: &db}

		// create tables
		if *dbActions == "create" {
			d.createTables()
		}
		// drop tables
		if *dbActions == "drop" {
			d.dropTables()
		}
		return
	}
	r := render.New(render.Options{Layout: "layout"})
	h := DBHandler{db: &db, r: r}

	mux := bone.New()
	mux.Get("/", http.HandlerFunc(homeHandler))
	mux.Post("/stubos", http.HandlerFunc(h.stubosCreateHandler))
	mux.Get("/stubos", http.HandlerFunc(h.stuboShowHandler))
	mux.Delete("/stubos/:id", http.HandlerFunc(h.stuboDestroyHandler))
	mux.Get("/stubos/:id", http.HandlerFunc(h.stuboDetailedHandler))
	mux.Get("/stubos/:id/scenarios/:scenario", http.HandlerFunc(h.scenarioDetailedHandler))
	// handling static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	n := negroni.Classic()
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(mux)

	log.WithFields(log.Fields{
		"port": port,
	}).Info("Overseer is starting")

	n.Run(*port)
}
