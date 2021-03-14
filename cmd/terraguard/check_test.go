package main

import (
	"errors"
	"os/exec"
	"strings"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		description string
		args        []string
		outputs     []string
		err         error
	}{{
		description: "when passed an unknown '--foo' flag",
		args:        []string{"--foo"},
		outputs:     []string{"Error: unknown flag: --foo\n"},
		err:         errors.New("exit status 1"),
	}, {
		description: "when passed an unknown 'foo' flag",
		args:        []string{"foo"},
		outputs: []string{
			"Error: required flag(s) \"plan\", \"resources\" not set",
		},
		err: errors.New("exit status 1"),
	}, {
		description: "when passed '--help'",
		args:        []string{"--help"},
		outputs: []string{
			"Check if a Terraform plan seeks to modify the specified resources",
			"Usage:",
			"Flags:",
		},
		err: nil,
	}, {
		description: "when passed '-h'",
		args:        []string{"-h"},
		outputs: []string{
			"Check if a Terraform plan seeks to modify the specified resources",
			"Usage:",
			"Flags:",
		},
		err: nil,
	}, {
		description: "when passed no arguments",
		args:        []string{""},
		outputs: []string{
			"Error: required flag(s) \"plan\", \"resources\" not set",
		},
		err: errors.New("exit status 1"),
	}, {
		description: "when the plan JSON file does not exist",
		args: []string{
			"--plan=does_not_exist.json",
			"--resources=foo",
		},
		outputs: []string{
			"Error: open does_not_exist.json: no such file or directory",
		},
		err: errors.New("exit status 1"),
	}, {
		description: "when the plan JSON does not indicate changes to the specified resource",
		args: []string{
			"--plan=../../test_fixtures/basic_plan.json",
			"--resources=foo",
		},
		outputs: []string{
			"",
		},
		err: nil,
	}}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			args := []string{}
			args = append([]string{"check"}, test.args...)

			output, err := exec.Command("./terraguard", args...).CombinedOutput()

			if test.err == nil && err != nil {
				t.Errorf("expected '%s' not to error; got '%v'", args, err)
			}

			if test.err != nil && err == nil {
				t.Errorf("expected '%s' to error with '%s', but it didn't error", args, test.err.Error())
			}

			if test.err != nil && err != nil && test.err.Error() != err.Error() {
				t.Errorf("expected '%s' to error with '%s'; got '%s'", args, test.err.Error(), err.Error())
			}

			for _, o := range test.outputs {
				if !strings.Contains(string(output), o) {
					t.Errorf("expected '%s' to include output '%s'; got '%s'", args, o, output)
				}
			}
		})
	}
}
