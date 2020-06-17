// the objective is to run nightly builds for black box testing for the entire
// chain
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var basescenarioPath = "scenarios"

type Scenario struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	call `json:"call"`
}

type call struct {
	ChainID  string `json:"chainId"`
	HTTPCall `json:"httpCall"`
}

type HTTPCall struct {
	Endpoint    string      `json:"endPoint"`
	Method      string      `json:"method"`
	ContentType string      `json:"contentType"`
	Headers     http.Header `json:"headers"`
	Body        string      `json:"body"`
}

func (h HTTPCall) Call() (*http.Response, error) {
	switch h.Method {
	case http.MethodGet:
		return http.Get(h.Endpoint)
	case http.MethodPost:
		return http.Post(h.Endpoint, h.ContentType, bytes.NewBufferString(h.Body))
	}
	return nil, errors.New("Invalid method")
}

func main() {
	// parse scenarios
	fs, err := ioutil.ReadDir(baseScenarioPath)
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		absFp, err := filepath.Abs(baseScenarioPath + "/" + f.Name())
		if err != nil {
			panic(err)
		}
		fd, err := ioutil.ReadFile(absFp)
		if err != nil {
			panic(err)
		}
		var s Scenario
		err = json.Unmarshal(fd, &s)
		if err != nil {
			panic(err)
		}
		fmt.Println(s.Call())
	}
}
