package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
)

// DBHandler used for passing database connection to handlers
type DBHandler struct {
	db *gorm.DB
	r  *render.Render
}

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

	// response := []byte("Hello!")
	// w.Write(response)
	ren.HTML(w, http.StatusOK, "example", nil)

}

func (h *DBHandler) stuboShowHandler(rw http.ResponseWriter, req *http.Request) {
	var stubos []Stubo
	h.db.Find(&stubos)
	if stubos == nil {
		h.r.JSON(rw, http.StatusOK, "[]")
	} else {
		h.r.JSON(rw, http.StatusOK, &stubos)
	}
}

func (h *DBHandler) scenarioDetailedHandler(rw http.ResponseWriter, req *http.Request) {
	// stubo ID, should be stored in database
	id := bone.GetValue(req, "id")
	scenario := bone.GetValue(req, "scenario")

	rw.Write([]byte("id:" + string(id) + ", scenario: " + scenario))
}
