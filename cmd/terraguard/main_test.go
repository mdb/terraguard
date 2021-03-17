package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	// compile a 'terraguard' for for use in running tests
	exe := exec.Command("go", "build", "-ldflags", "-X main.version=test", "-o", "terraguard")
	err := exe.Run()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
