package main

import (
	"errors"
	"os/exec"
	"strings"
	"testing"

	"github.com/mdb/terraguard/cmd"
)

func TestCheck(t *testing.T) {
	exitErr := errors.New("exit status 1")

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
			"Error: required flag(s) \"guard\", \"plan\" not set",
		},
		err: exitErr,
	}, {
		description: "when passed '--help'",
		args:        []string{"--help"},
		outputs: []string{
			cmd.CheckLongDescription,
			"Usage:",
			"Flags:",
		},
		err: nil,
	}, {
		description: "when passed '-h'",
		args:        []string{"-h"},
		outputs: []string{
			cmd.CheckLongDescription,
			"Usage:",
			"Flags:",
		},
		err: nil,
	}, {
		description: "when passed no arguments",
		args:        []string{""},
		outputs: []string{
			"Error: required flag(s) \"guard\", \"plan\" not set",
		},
		err: errors.New("exit status 1"),
	}, {
		description: "when the plan JSON file does not exist",
		args: []string{
			"--plan=does_not_exist.json",
			"--guard=foo",
		},
		outputs: []string{
			"Error: open does_not_exist.json: no such file or directory",
		},
		err: exitErr,
	}, {
		description: "when the plan JSON does not indicate changes to the specified resource",
		args: []string{
			"--plan=../../testdata/show-plan.json",
			"--guard=local_file.greeting_no_change",
		},
		outputs: []string{
			"",
		},
		err: nil,
	}, {
		description: "when the plan JSON indicates changes to the specified resource",
		args: []string{
			"--plan=../../testdata/show-plan.json",
			"--guard=local_file.greeting_one",
		},
		outputs: []string{
			"Error: ../../testdata/show-plan.json indicates changes to guarded resources:\n\nlocal_file.greeting_one",
		},
		err: exitErr,
	}, {
		description: "when the plan JSON indicates changes to the specified '*'-prefixed resource",
		args: []string{
			"--plan=../../testdata/show-plan.json",
			"--guard=*file.greeting_one",
		},
		outputs: []string{
			"Error: ../../testdata/show-plan.json indicates changes to guarded resources:\n\nlocal_file.greeting_one",
		},
		err: exitErr,
	}, {
		description: "when the plan JSON indicates changes to the specified '*'-suffixed resource",
		args: []string{
			"--plan=../../testdata/show-plan.json",
			"--guard=local_file.greeting_*",
		},
		outputs: []string{
			"Error: ../../testdata/show-plan.json indicates changes to guarded resources:\n\nlocal_file.greeting_one\nlocal_file.greeting_two",
		},
		err: exitErr,
	}, {
		description: "when the plan JSON indicates changes to the specified '*'-prefixed and '*'-suffixed resource",
		args: []string{
			"--plan=../../testdata/show-plan.json",
			"--guard=*local*",
		},
		outputs: []string{
			"Error: ../../testdata/show-plan.json indicates changes to guarded resources:\n\nlocal_file.greeting_one\nlocal_file.greeting_two",
		},
		err: exitErr,
	}, {
		description: "when the plan JSON indicates changes to one of multiple comma-separated resources",
		args: []string{
			"--plan=../../testdata/show-plan.json",
			"--guard=local_file.greeting_one, bar",
		},
		outputs: []string{
			"Error: ../../testdata/show-plan.json indicates changes to guarded resources:\n\nlocal_file.greeting_one",
		},
		err: exitErr,
	}, {
		description: "when the plan JSON indicates changes to multiple comma-separated resources",
		args: []string{
			"--plan=../../testdata/show-plan.json",
			"--guard=local_file.greeting_one, local_file.greeting_two",
		},
		outputs: []string{
			"Error: ../../testdata/show-plan.json indicates changes to guarded resources:\n\nlocal_file.greeting_one\nlocal_file.greeting_two",
		},
		err: exitErr,
	}, {
		description: "when the plan JSON indicates changes to multiple individually specified resources",
		args: []string{
			"--plan=../../testdata/show-plan.json",
			"--guard=local_file.greeting_one",
			"--guard=local_file.greeting_two",
		},
		outputs: []string{
			"Error: ../../testdata/show-plan.json indicates changes to guarded resources:\n\nlocal_file.greeting_one\nlocal_file.greeting_two",
		},
		err: exitErr,
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
