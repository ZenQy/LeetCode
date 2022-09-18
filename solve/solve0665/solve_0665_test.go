package solve0665

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPossibility(t *testing.T) {
	assert.Equal(t, checkPossibility([]int{1, 2, 3, 4}), true)
	assert.Equal(t, checkPossibility([]int{4, 1, 2, 3}), true)
	assert.Equal(t, checkPossibility([]int{2, 3, 4, 1}), true)
	assert.Equal(t, checkPossibility([]int{2, 3, 1, 4}), true)
	assert.Equal(t, checkPossibility([]int{1, 2, 4, 3}), true)
}
