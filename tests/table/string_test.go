package strings

import (
	"strings"
	"testing"
)

const dftMsg = "%s (part: %s) - Index: expected (%d) <> received (%d)"

func TestIndex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		text   string
		part   string
		expect int
	}{
		{"Cod3r is nice", "Cod3r", 0},
		{"", "", 0},
		{"Ops", "ops", -1},
		{"Andre", "e", 4},
	}

	for _, test := range tests {
		t.Logf("Mass: %v", test)

		index := strings.Index(test.text, test.part)

		if index != test.expect {
			t.Errorf(dftMsg, test.text, test.part, test.expect, index)
		}
	}
}
