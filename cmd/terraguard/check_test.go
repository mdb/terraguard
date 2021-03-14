package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		arg     string
		outputs []string
		err     error
	}{{
		arg:     "--foo",
		outputs: []string{"Error: unknown flag: --foo\n"},
		err:     errors.New("exit status 1"),
	}, {
		arg: "foo",
		outputs: []string{
			"Error: required flag(s) \"plan\", \"resources\" not set",
		},
		err: errors.New("exit status 1"),
	}, {
		arg: "--help",
		outputs: []string{
			"Check if a Terraform plan seeks to modify the specified resources",
			"Usage:",
			"Flags:",
		},
		err: nil,
	}, {
		arg: "-h",
		outputs: []string{
			"Check if a Terraform plan seeks to modify the specified resources",
			"Usage:",
			"Flags:",
		},
		err: nil,
	}, {
		arg: "",
		outputs: []string{
			"Error: required flag(s) \"plan\", \"resources\" not set",
		},
		err: errors.New("exit status 1"),
	}}

	for _, test := range tests {
		t.Run(fmt.Sprintf("when terraguard is passed '%s'", test.arg), func(t *testing.T) {
			output, err := exec.Command("./terraguard", "check", test.arg).CombinedOutput()

			if test.err == nil && err != nil {
				t.Errorf("expected '%s' not to error; got '%v'", test.arg, err)
			}

			if test.err != nil && err == nil {
				t.Errorf("expected '%s' to error with '%s', but it didn't error", test.arg, test.err.Error())
			}

			if test.err != nil && err != nil && test.err.Error() != err.Error() {
				t.Errorf("expected '%s' to error with '%s'; got '%s'", test.arg, test.err.Error(), err.Error())
			}

			for _, o := range test.outputs {
				if !strings.Contains(string(output), o) {
					t.Errorf("expected '%s' to include output '%s'; got '%s'", test.arg, o, output)
				}
			}
		})
	}
}
