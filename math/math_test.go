package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		t.Errorf("FAIL: wanted 11 but received %d", sum)
	}
}
