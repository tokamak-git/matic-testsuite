package blackbox

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Senarios(t *testing.T) {
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
		assert.Equal(t, s.Out, "")
	}
}
