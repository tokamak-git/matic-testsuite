// the objective is to run nightly builds for black box testing for the entire
// chain
package blackbox

import (
	"github.com/maticnetwork/matic-testsuite/caller"
)

var baseScenarioPath = "scenarios"

type Scenario struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	caller.Call `json:"call"`
	Out         interface{} `json:"output"`
}
