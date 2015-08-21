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
}

// getScenariosDetail gets and returns all scenarios with details
func (s *Stubo) getScenariosDetail() ([]byte, error) {
	path := "stubo/api/v2/scenarios/detail"
	fullPath := fmt.Sprintf("%s://%s:%s/%s", s.protocol, s.host, s.port, path)
	fmt.Println(fullPath)
	return []byte(""), nil
}
