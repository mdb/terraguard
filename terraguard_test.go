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
		name:               "when checking changes to a non-existing module address name",
		guardedChanges:     []string{"foo"},
		expectedViolations: 0,
		planJSON:           "basic_plan.json",
	}, {
		name:               "when checking changes to an existing module address name",
		guardedChanges:     []string{"data.null_data_source.baz"},
		expectedViolations: 1,
		planJSON:           "basic_plan.json",
	}, {
		name:               "when checking changes to an existing module address name suffixed with a '*'",
		guardedChanges:     []string{"null_resource.baz*"},
		expectedViolations: 3,
		planJSON:           "basic_plan.json",
	}, {
		name:               "when checking changes to an existing module address name prefixed with a '*'",
		guardedChanges:     []string{"*.foo.null_resource.foo"},
		expectedViolations: 1,
		planJSON:           "basic_plan.json",
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g, err := NewGuard("test_fixtures/"+test.planJSON, test.guardedChanges)
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
