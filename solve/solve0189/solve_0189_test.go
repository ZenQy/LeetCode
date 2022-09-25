package solve0189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums1, 3)
	assert.Equal(t, nums1, []int{5, 6, 7, 1, 2, 3, 4})
	nums2 := []int{-1, -100, 3, 99}
	rotate(nums2, 2)
	assert.Equal(t, nums2, []int{3, 99, -1, -100})
}
