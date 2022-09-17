package solve0485

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	assert.Equal(t, findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}), 3)
	assert.Equal(t, findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}), 2)
}
