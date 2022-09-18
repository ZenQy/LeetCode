package solve0453

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMoves(t *testing.T) {
	assert.Equal(t, minMoves([]int{1, 2, 3}), 3)
	assert.Equal(t, minMoves([]int{1, 1, 1}), 0)
}
