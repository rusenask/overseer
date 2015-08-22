package main

import (
	"fmt"
	"net/http"
)

// Stubo structure to be injected into functions to perform HTTP calls
type Stubo struct {
	HTTPClient *http.Client
	host       string
	port       string
	protocol   string
	uri        string
}

type scenarioDoc struct {
	id   string `bson:"_id"`
	name string `bson:"name"`
// Model provides a default model struct, you could embed it in your struct
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
}

// Scenario object
type Scenario struct {
	name string
	Stubo
}

func (s *Scenario) getStubs() ([]string, error) {
	path := s.Stubo.uri + "/stubo/api/v2/scenarios/objects" + s.name + "/stubs"
	fmt.Println(path)
	return []string{"nope", "nope2"}, nil
}

// getScenariosDetail gets and returns all scenarios with details
func (s *Stubo) getScenariosDetail() ([]byte, error) {
	path := "stubo/api/v2/scenarios/detail"
	fullPath := fmt.Sprintf("%s://%s:%s/%s", s.protocol, s.host, s.port, path)
	fmt.Println(fullPath)
	return []byte(""), nil
}
