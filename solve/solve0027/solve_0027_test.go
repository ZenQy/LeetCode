package solve0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	assert.Equal(t, []int{2, 2}, nums[:removeElement(nums, val)])
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	assert.Equal(t, []int{0, 1, 4, 0, 3}, nums[:removeElement(nums, val)])
}
