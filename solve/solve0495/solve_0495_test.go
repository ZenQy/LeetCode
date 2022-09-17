package solve0495

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPoisonedDuration(t *testing.T) {
	assert.Equal(t, findPoisonedDuration([]int{1, 4}, 2), 4)
	assert.Equal(t, findPoisonedDuration([]int{1, 2}, 2), 3)
}
