package cui_test

import (
	"testing"

	"github.com/erdaltsksn/cui"
)

func TestVersionCmd(t *testing.T) {
	if err := cui.VersionCmd.Execute(); err != nil {
		t.Error("Got:", err, "Want:", nil)
	}
}
