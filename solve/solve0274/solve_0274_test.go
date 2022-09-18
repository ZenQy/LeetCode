package solve0274

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHIndex(t *testing.T) {
	assert.Equal(t, hIndex([]int{3, 0, 6, 1, 5}), 3)
	assert.Equal(t, hIndex([]int{1, 3, 1}), 1)
}
