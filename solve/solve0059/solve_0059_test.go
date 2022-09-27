package solve0059

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	assert.Equal(t, generateMatrix(3), [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}})
}
