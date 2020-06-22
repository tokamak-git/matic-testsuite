package asserter

import (
	"testing"

	"github.com/maticnetwork/matic-testsuite/caller"
)

type HTTPAssert struct {
	caller.HTTPCall
	Assert
}

func (h HTTPAssert) Exec(t testing.T) error {
	resp, err := h.Call()
	// TODO need to rethink this, may required to check errors as well
	if err != nil {
		return err
	}

	equal, err := h.Assert.Check(t, resp)
	if equal == false {
		t.Error("Failed test")
	}
}
