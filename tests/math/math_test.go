package math

import "testing"

const defaultErrorMsg = "Expected %v, received %v"

func TestAvg(t *testing.T) {
	t.Parallel()
	expect := 7.28
	value := Avg(7.2, 9.9, 6.1, 5.9)

	if value != expect {
		t.Errorf(defaultErrorMsg, expect, value)
	}
}
