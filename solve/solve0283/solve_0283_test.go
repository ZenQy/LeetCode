package solve0283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	assert.Equal(t, moveZeroes([]int{0, 1, 0, 3, 12}), []int{1, 3, 12, 0, 0})
	assert.Equal(t, moveZeroes([]int{0}), []int{0})
}
