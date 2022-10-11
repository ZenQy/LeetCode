package solve0303

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArr := Constructor(nums)
	assert.Equal(t, numArr.SumRange(0, 2), 1)
	assert.Equal(t, numArr.SumRange(2, 5), -1)
	assert.Equal(t, numArr.SumRange(0, 5), -3)
}
