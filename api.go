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

// Session stores session data (can be found in scenario details)
type Session struct {
	Status   string `json:"status"`
	Loaded   string `json:"loaded"`
	Name     string `json:"name"`
	LastUsed string `json:"last_used"`
}

// Scenario structure for gettting scenario objects (not detailed). This struct
// can be used for both scenario list and scenario detailed list
type Scenario struct {
	Name      string    `json:"name"`
	Ref       string    `json:"ScenarioRef"`
	SpaceUsed int       `json:"space_used_kb"`
	Sessions  []Session `json:"sessions"`
	Recorded  string    `json:"recorded"`
	StubCount int       `json:"stub_count"`
}

// ScenarioResponse structure for unmarshaling JSON structures from API v2
type ScenarioResponse struct {
	Data []Scenario `json:"data"`
}

// RequestContains is an array of strings tha define how stub should be matched
type RequestContains struct {
	Contains []string `json:"contains"`
}

// StubRequest stores information about request
type StubRequest struct {
	BodyPatterns RequestContains `json:"bodyPatterns"`
	Method       string          `json:"method"`
}

// StubArgs stores information about arguments such as session
type StubArgs struct {
	Priority string `json:"priority"`
	Session  string `json:"session"`
}

// ResponseFromStubo stores information about response that stubo should return
type ResponseFromStubo struct {
	StatusCode int      `json:"status"`
	Body       []string `json:"body"`
}

// StubDetails stores information about the stub itself
type StubDetails struct {
	Priority int               `json:"priority"`
	Request  StubRequest       `json:"request"`
	Args     StubArgs          `json:"args"`
	Response ResponseFromStubo `json:"response"`
}

// Stub - stub doc
type Stub struct {
	Stub      StubDetails `json:"stub"`
	SpaceUsed int         `json:"space_used"`
	Recorded  string      `json:"recorded"`
	Matcher   string      `json:"matchers_hash"`
}

// StuboResponse - JSON response from stubo API
type StuboResponse struct {
	Data    []Stub `json:"data"`
	Version string `json:"version"`
}

// getScenarios gets and returns all scenarios with details
func (c *Client) getScenarios(uri string) ([]Scenario, error) {
	path := "/stubo/api/v2/scenarios/detail"
	fullPath := uri + path
	log.WithFields(log.Fields{
		"name":          "",
		"urlPath":       path,
		"headers":       "",
		"body":          "",
		"requestMethod": "GET",
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

func (c *Client) getScenarioDetails(uri, scenario string) (Scenario, error) {
	fullPath := uri + "/stubo/api/v2/scenarios/objects/" + scenario

	// logging
	log.WithFields(log.Fields{
		"name":          scenario,
		"urlPath":       fullPath,
		"headers":       "",
		"body":          "",
		"requestMethod": "GET",
	}).Debug("Getting scenario details")

	respBody, err := c.GetResponseBody(fullPath)

	if err != nil {
		return Scenario{}, err
	}

	var data Scenario
	err = json.Unmarshal(respBody, &data)

	if err != nil {
		return Scenario{}, err
	}
	return data, nil
}

func (c *Client) getScenarioStubs(uri, scenario string) ([]Stub, error) {
	fullPath := uri + "/stubo/api/v2/scenarios/objects/" + scenario + "/stubs"

	// logging
	log.WithFields(log.Fields{
		"name":          scenario,
		"urlPath":       fullPath,
		"headers":       "",
		"body":          "",
		"requestMethod": "GET",
	}).Debug("Getting scenario stubs")

	respBody, err := c.GetResponseBody(fullPath)

	if err != nil {
		return []Stub{}, err
	}

	var data StuboResponse
	err = json.Unmarshal(respBody, &data)

	if err != nil {
		return []Stub{}, err
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
