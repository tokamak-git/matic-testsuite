package senarios

import "matic-testsuite/runner"

var Senarios []Instance

type Instance struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Setup     []runner.Action
	Case      []runner.Action
	Assertion []runner.Action
}
