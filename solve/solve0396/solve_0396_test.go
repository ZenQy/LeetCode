package solve0396

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxRotateFunction(t *testing.T) {
	assert.Equal(t, maxRotateFunction([]int{4, 3, 2, 6}), 26)
	assert.Equal(t, maxRotateFunction([]int{100}), 0)
}
