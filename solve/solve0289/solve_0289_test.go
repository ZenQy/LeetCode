package solve0289

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameOfLife(t *testing.T) {
	board1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	board2 := [][]int{{1, 1}, {1, 0}}
	gameOfLife(board1)
	gameOfLife(board2)
	assert.Equal(t, board1, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}})
	assert.Equal(t, board2, [][]int{{1, 1}, {1, 1}})
}
