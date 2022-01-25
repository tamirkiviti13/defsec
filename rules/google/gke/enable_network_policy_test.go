package gke

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/google/gke"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnableNetworkPolicy(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    gke.GKE
		expected bool
	}{
		{
			name:     "positive result",
			input:    gke.GKE{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    gke.GKE{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Google.GKE = test.input
			results := CheckEnableNetworkPolicy.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Rule().LongID() == CheckEnableNetworkPolicy.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}