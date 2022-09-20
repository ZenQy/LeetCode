package solve0419

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBattleships(t *testing.T) {
	assert.Equal(t, countBattleships([][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}), 2)
	assert.Equal(t, countBattleships([][]byte{{'.'}}), 0)
}
