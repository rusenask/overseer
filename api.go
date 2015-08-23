package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// Client structure to be injected into functions to perform HTTP calls
type Client struct {
	HTTPClient *http.Client
}

// Scenario structure for gettting scenario objects (not detailed)
type Scenario struct {
	Name string `json:"name"`
	Ref  string `json:"ScenarioRef"`
}

// ScenarioResponse structure for unmarshaling JSON structures from API v2
type ScenarioResponse struct {
	Data []Scenario `json:"data"`
}

// getScenarios gets and returns all scenarios with details
func (c *Client) getScenarios(uri string) ([]Scenario, error) {
	path := "/stubo/api/v2/scenarios"
	fullPath := uri + path
	log.WithFields(log.Fields{
		"name":          "",
		"urlPath":       path,
		"headers":       "",
		"body":          "",
		"requestMethod": "",
	}).Debug("Getting scenarios")
	respBody, err := c.GetResponseBody(fullPath)

	if err != nil {
		return []Scenario{}, err
	}

	var data ScenarioResponse
	err = json.Unmarshal(respBody, &data)

	if err != nil {
		return []Scenario{}, err
	}

	return data.Data, nil
}

// GetResponseBody calls stubo
func (c *Client) GetResponseBody(uri string) ([]byte, error) {
	// logging get Response body
	log.WithFields(log.Fields{
		"uri": uri,
	}).Info("Getting response body")
	resp, err := c.HTTPClient.Get(uri)

	if err != nil {
		// logging get error
		log.WithFields(log.Fields{
			"error": err.Error(),
			"uri":   uri,
		}).Warn("Failed to get response from Stubo!")

		return []byte(""), err
	}
	defer resp.Body.Close()
	// reading resposne body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// logging read error
		log.WithFields(log.Fields{
			"error": err.Error(),
			"uri":   uri,
		}).Warn("Failed to read response from Stubo!")

		return []byte(""), err
	}
	return body, nil
}
