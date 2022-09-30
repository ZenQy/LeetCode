package solve0073

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	mat1 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	mat2 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(mat1)
	setZeroes(mat2)
	assert.Equal(t, mat1, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}})
	assert.Equal(t, mat2, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}})
}
