package solve0283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums1)
	assert.Equal(t, nums1, []int{1, 3, 12, 0, 0})
	nums2 := []int{0}
	moveZeroes(nums2)
	assert.Equal(t, nums2, []int{0})
}
