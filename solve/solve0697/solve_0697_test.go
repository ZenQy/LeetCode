package solve0697

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestSubArray(t *testing.T) {
	assert.Equal(t, findShortestSubArray([]int{1, 2, 3}), 1)
	assert.Equal(t, findShortestSubArray([]int{1, 2, 2, 3, 1}), 2)
	assert.Equal(t, findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}), 6)
}
