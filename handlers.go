package main

import (
	"fmt"
	"net/http"
	"strconv"

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

	ren.HTML(w, http.StatusOK, "example", nil)
}

func (h *DBHandler) stuboShowHandler(rw http.ResponseWriter, req *http.Request) {
	var stuboInstances []Stubo
	// Getting all stubo instances
	stuboInstances = h.getAllInstances()

	// setting logger
	log.WithFields(log.Fields{
		"count": len(stuboInstances),
	}).Info("Getting all Stubo instances")

	newmap := map[string]interface{}{"metatitle": "Stubo Instances", "Instances": stuboInstances}
	h.r.HTML(rw, http.StatusOK, "stubos", newmap)
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

	// validate form
	if err := binding.Bind(req, stuboForm); err.Handle(rw) {
		fmt.Println(err.Error())
		return
	}
	// assing form variables to stubo struct
	stubo := Stubo{Name: stuboForm.Name, Version: stuboForm.Version, Hostname: stuboForm.Hostname,
		Port: stuboForm.Port, Protocol: stuboForm.Protocol}

	h.db.Create(&stubo)
	log.WithFields(log.Fields{
		"id":       id,
		"url_path": req.URL.Path,
	}).Info("Stubo added")

	// TODO: should add some messages to user about successful creation
	h.stuboShowHandler(rw, req)
}

// stuboDestroyHandler deletes a stubo instance from the db.
func (h *DBHandler) stuboDestroyHandler(rw http.ResponseWriter, req *http.Request) {
	id := bone.GetValue(req, "id")
	stubo := Stubo{}
	h.db.Delete(&stubo, id)

	log.WithFields(log.Fields{
		"id":       id,
		"url_path": req.URL.Path,
	}).Info("Stubo deleted")

	h.r.JSON(rw, http.StatusOK, map[string]string{"data": "Stubo instance deleted!"})
}

func (h *DBHandler) stuboDetailedHandler(rw http.ResponseWriter, req *http.Request) {
	id := bone.GetValue(req, "id")
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(rw, "Bad stubo ID.", 400)
	}
	var stubo Stubo

	stubo = h.getStuboDetails(u)
	stuboURI := stubo.Protocol + "://" + stubo.Hostname + ":" + stubo.Port
	client := &Client{&http.Client{}}
	scenarios, err := client.getScenarios(stuboURI)
	if err != nil {
		http.Error(rw, "Failed to get scenarios from Stubo!", 400)
	}
	log.WithFields(log.Fields{
		"id":             id,
		"url_path":       req.URL.Path,
		"scenario_count": len(scenarios),
	}).Info("Stubo details fetched")

	newmap := map[string]interface{}{"metatitle": "Stubo Details", "Scenarios": scenarios}
	h.r.HTML(rw, http.StatusOK, "stuboDetails", newmap)
}

func (h *DBHandler) scenarioDetailedHandler(rw http.ResponseWriter, req *http.Request) {
	// stubo ID, should be stored in database
	id := bone.GetValue(req, "id")
	scenario := bone.GetValue(req, "scenario")

	rw.Write([]byte("id:" + string(id) + ", scenario: " + scenario))
}
