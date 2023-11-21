package baz_test

import (
	"testing"

	"github.com/cybozu-private/test-go-test/foo/utils"
)

const pkgName = "foo/baz"

func TestBaz(t *testing.T) {
	utils.TimingTest(t, pkgName)
}
