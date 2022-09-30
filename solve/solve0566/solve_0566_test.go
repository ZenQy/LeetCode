package solve0566

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixReshape(t *testing.T) {
	assert.Equal(t, matrixReshape([][]int{{1, 2, 3, 4}}, 2, 2), [][]int{{1, 2}, {3, 4}})
	assert.Equal(t, matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4), [][]int{{1, 2, 3, 4}})
	assert.Equal(t, matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4), [][]int{{1, 2}, {3, 4}})
}
