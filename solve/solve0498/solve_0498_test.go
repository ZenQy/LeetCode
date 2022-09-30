package solve0498

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDiagonalOrder(t *testing.T) {
	assert.Equal(t, findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
		[]int{1, 2, 4, 7, 5, 3, 6, 8, 9})
	assert.Equal(t, findDiagonalOrder([][]int{{1, 2}, {3, 4}}),
		[]int{1, 2, 3, 4})
	assert.Equal(t, findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}}),
		[]int{1, 2, 4, 5, 3, 6})
	assert.Equal(t, findDiagonalOrder([][]int{{1}}),
		[]int{1})
}
