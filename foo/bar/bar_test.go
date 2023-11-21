package bar_test

import (
	"testing"

	"github.com/cybozu-private/test-go-test/foo/utils"
)

const pkgName = "foo/bar"

func TestBar(t *testing.T) {
	utils.TimingTest(t, pkgName)
}
