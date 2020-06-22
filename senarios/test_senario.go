package senarios

import (
	"matic-testsuite/runner"
	"net/http"

	"github.com/maticnetwork/matic-testsuite/caller"
)

type TestSenario struct {
	caller.HTTPCall
}

func init() {
	t := TestSenario{
		HTTPCall: caller.HTTPCall{Endpoint: "http://localhost:8080", Method: http.MethodPost, ContentType: "text", Body: "test Body"},
	}
	Senarios = append(Senarios, S{Case: []runner.Action{t}})
}

func (t TestSenario) Exec() error {
	_, err := t.HTTPCall.Call()
	return err
}
