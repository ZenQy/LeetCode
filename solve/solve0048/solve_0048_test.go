package solve0048

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	mat1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	mat2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(mat1)
	rotate(mat2)
	assert.Equal(t, mat1, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}})
	assert.Equal(t, mat2, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}})
}
