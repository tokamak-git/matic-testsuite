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

var baseSenarioPath = "senarios"

type senario struct {
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
	case "POST":
		return http.Post(h.Endpoint, h.ContentType, bytes.NewBufferString(h.Body))
	}
	return nil, errors.New("Invalid method")
}

func main() {
	// parse senarios
	fs, err := ioutil.ReadDir(baseSenarioPath)
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		absFp, err := filepath.Abs(baseSenarioPath + "/" + f.Name())
		if err != nil {
			panic(err)
		}
		fd, err := ioutil.ReadFile(absFp)
		if err != nil {
			panic(err)
		}
		var s senario
		err = json.Unmarshal(fd, &s)
		if err != nil {
			panic(err)
		}
		fmt.Println(s.Call())
	}
}
