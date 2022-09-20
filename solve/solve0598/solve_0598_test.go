package solve0598

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCount(t *testing.T) {
	assert.Equal(t, maxCount(3, 3, [][]int{{2, 2}, {3, 3}}), 4)
	assert.Equal(t, maxCount(18, 3, [][]int{{16, 1}, {14, 3}, {14, 2}, {4, 1}, {10, 1}, {11, 1}, {8, 3}, {16, 2}, {13, 1}, {8, 3}, {2, 2}, {9, 1}, {3, 1}, {2, 2}, {6, 3}}), 2)
	assert.Equal(t, maxCount(4, 3, [][]int{{2, 2}, {1, 3}, {2, 3}, {1, 2}, {2, 2}}), 2)
	assert.Equal(t, maxCount(3, 3, [][]int{}), 9)
}
