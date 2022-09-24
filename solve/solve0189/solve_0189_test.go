package solve0189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	assert.Equal(t, rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3), []int{5, 6, 7, 1, 2, 3, 4})
	assert.Equal(t, rotate([]int{-1, -100, 3, 99}, 2), []int{3, 99, -1, -100})
}
