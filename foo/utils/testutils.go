package utils

import (
	"fmt"
	"testing"
	"time"
)

func TimingTest(t *testing.T, pkgName string) {
	was := time.Now()
	fmt.Println("Pre sleep in", pkgName, was)
	// sleep
	time.Sleep(5 * time.Second)
	fmt.Println("Post sleep in", pkgName, time.Now())

	if time.Since(was)-5*time.Second > time.Second {
		t.Fatalf("sleep was too long in %s, actual sleep was %v\n", pkgName, time.Since(was))
	}
}
