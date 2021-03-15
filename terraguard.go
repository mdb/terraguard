package terraguard

import (
	"encoding/json"
	"os"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
)

// Guard is used to inspect a Terraform plan JSON
// for changes to a list of guarded resource names.
type Guard struct {
	// Plan is a a *tfjson.Plan
	Plan *tfjson.Plan

	// GuardedResources is a list of guarded Terraform
	// resource names.
	GuardedResources []string
}

// Check checks a Guard.Plan for changes to guarded resources.
func (g *Guard) Check() (bool, []*tfjson.ResourceChange) {
	hasGuardedChanges := false
	guardedChanges := []*tfjson.ResourceChange{}

	for _, rc := range g.Plan.ResourceChanges {
		for _, gr := range g.GuardedResources {
			if g.isGuarded(rc.Address, gr) {
				guardedChanges = append(guardedChanges, rc)
			}
		}
	}

	if len(guardedChanges) > 0 {
		hasGuardedChanges = true
	}

	return hasGuardedChanges, guardedChanges
}

func (g *Guard) isGuarded(address, guardedAddress string) bool {
	sanitizedAddress := strings.ReplaceAll(address, " ", "")
	sanitizedGuardedAddress := strings.ReplaceAll(guardedAddress, " ", "")

	if strings.Contains(sanitizedGuardedAddress, "*") {
		baseAddressName := strings.ReplaceAll(sanitizedGuardedAddress, "*", "")

		return strings.Contains(sanitizedAddress, baseAddressName)
	}

	return sanitizedAddress == sanitizedGuardedAddress
}

// NewGuard returns a new Guard built from the Terraform
// plan JSON file and guarded Terraform resource address
// names it is passed.
func NewGuard(planFile string, guardedResources []string) (*Guard, error) {
	g := &Guard{
		GuardedResources: guardedResources,
	}

	f, err := os.Open(planFile)
	if err != nil {
		return g, err
	}
	defer f.Close()

	var plan *tfjson.Plan
	if err := json.NewDecoder(f).Decode(&plan); err != nil {
		return g, err
	}

	g.Plan = plan

	return g, nil
}
