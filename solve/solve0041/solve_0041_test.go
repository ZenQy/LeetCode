package solve0041

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	assert.Equal(t, firstMissingPositive([]int{1, 1}), 2)
	assert.Equal(t, firstMissingPositive([]int{1, 2, 3}), 4)
	assert.Equal(t, firstMissingPositive([]int{3, 4, -1, 1}), 2)
	assert.Equal(t, firstMissingPositive([]int{7, 8, 9, 11, 12}), 1)
}
