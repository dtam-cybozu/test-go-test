package foo_test

import (
	"testing"

	"github.com/cybozu-private/test-go-test/foo/utils"
)

const pkgName = "foo"

func TestFoo(t *testing.T) {
	utils.TimingTest(t, pkgName)
}
