package solve0238

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	assert.Equal(t, productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6})
	assert.Equal(t, productExceptSelf([]int{-1, 1, 0, -3, 3}), []int{0, 0, 9, 0, 0})
}
