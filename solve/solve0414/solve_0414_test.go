package solve0414

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	assert.Equal(t, thirdMax([]int{3, 2, 1}), 1)
	assert.Equal(t, thirdMax([]int{1, 2, -2147483648}), -2147483648)
	assert.Equal(t, thirdMax([]int{1, 2}), 2)
	assert.Equal(t, thirdMax([]int{2, 2, 3, 1}), 1)

}
