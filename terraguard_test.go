package terraguard

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name               string
		guardedChanges     []string
		expectedViolations int
		planJSON           string
	}{{
		name:               "when checking changes to a non-existing resource address name",
		guardedChanges:     []string{"foo"},
		expectedViolations: 0,
		planJSON:           "testdata/show-plan.json",
	}, {
		name:               "when checking changes to an existing resource address name",
		guardedChanges:     []string{"local_file.greeting_one"},
		expectedViolations: 1,
		planJSON:           "testdata/show-plan.json",
	}, {
		name:               "when checking changes to an existing resource address name with NOOP changes",
		guardedChanges:     []string{"local_file.greeting_no_change"},
		expectedViolations: 0,
		planJSON:           "testdata/show-plan.json",
	}, {
		name:               "when checking changes to an existing resource address name suffixed with a '*'",
		guardedChanges:     []string{"local_file.greeting*"},
		expectedViolations: 2,
		planJSON:           "testdata/show-plan.json",
	}, {
		name:               "when checking changes to an existing address name prefixed with a '*'",
		guardedChanges:     []string{"*.greeting_one"},
		expectedViolations: 1,
		planJSON:           "testdata/show-plan.json",
	}, {
		name:               "when checking changes to an existing address name prefixed and suffixed with a '*'",
		guardedChanges:     []string{"*.greeting*"},
		expectedViolations: 2,
		planJSON:           "testdata/show-plan.json",
	}, {
		name:               "when checking changes to multiple existing address names with spaces",
		guardedChanges:     []string{"local_file.greeting_one", " local_file.greeting_two "},
		expectedViolations: 2,
		planJSON:           "testdata/show-plan.json",
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g, err := NewGuard(test.planJSON, test.guardedChanges)
			if err != nil {
				t.Errorf("expected NewGuard not to error; got '%v'", err)
			}

			found, violations := g.Check()

			expectedFound := false
			if test.expectedViolations > 0 {
				expectedFound = true
			}

			if found != expectedFound {
				t.Errorf("expected Check to return '%v'; got '%v'", expectedFound, found)
			}

			violationsLen := len(violations)
			if violationsLen != test.expectedViolations {
				t.Errorf("expected Check to return '%v' violations; got '%v'", test.expectedViolations, violationsLen)
			}
		})
	}
}
