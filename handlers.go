package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
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
	// var stuboInstances []Stubo
	var retData struct {
		Instances []Stubo
	}
	h.db.Find(&retData.Instances)
	fmt.Println(retData)

	newmap := map[string]interface{}{"metatitle": "created post", "Instances": retData.Instances}
	h.r.HTML(rw, http.StatusOK, "stubos", newmap)
	// if stbs == nil {
	// 	h.r.HTML(rw, http.StatusOK, "stubos", &stbs)
	// } else {
	// 	h.r.HTML(rw, http.StatusOK, "stubos", &stbs)
	// }
}

// stubosCreateHandler inserts a new guitar into the db.
func (h *DBHandler) stubosCreateHandler(rw http.ResponseWriter, req *http.Request) {
	h.stubosEdit(rw, req, 0)
}

// stubosEdit is shared between the create and update handler, since they share most of the logic.
func (h *DBHandler) stubosEdit(rw http.ResponseWriter, req *http.Request, id uint) {
	stuboForm := new(StuboForm)

	log.WithFields(log.Fields{
		"id":       id,
		"url_path": req.URL.Path,
	}).Info("Entering stubosEdti")

	if err := binding.Bind(req, stuboForm); err.Handle(rw) {
		return
	}
	stubo := Stubo{Name: stuboForm.Name, Version: stuboForm.Version, Hostname: stuboForm.Hostname,
		Port: stuboForm.Port, Protocol: stuboForm.Protocol}
	h.db.Create(&stubo)
	log.WithFields(log.Fields{
		"id":       id,
		"url_path": req.URL.Path,
	}).Info("Stubo added")
	h.r.HTML(rw, http.StatusOK, "stubos", &stubo)
}

func (h *DBHandler) scenarioDetailedHandler(rw http.ResponseWriter, req *http.Request) {
	// stubo ID, should be stored in database
	id := bone.GetValue(req, "id")
	scenario := bone.GetValue(req, "scenario")

	rw.Write([]byte("id:" + string(id) + ", scenario: " + scenario))
}
