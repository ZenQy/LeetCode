package solve0628

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumProduct(t *testing.T) {
	assert.Equal(t, maximumProduct([]int{-1, -2, -3, -4}), -6)
	assert.Equal(t, maximumProduct([]int{1, 2, 3, 4}), 24)
	assert.Equal(t, maximumProduct([]int{-1, 2, 3, 4}), 24)
}
