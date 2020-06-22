// package asserter is to be uses to check values and pass rules for the same
package asserter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Assert allows us to check values, once data is defined we can view it using
// the data block
// if specific logic is needed then we can pass a rule to the struct via rule
type Assert struct {
	Rule string
	Data []byte
}

func (a Assert) Check(t testing.T, i interface{}) (equal bool, err error) {
	// Rule can be used to
	equal = assert.Equal(t, a.Data, i)
	return
}
