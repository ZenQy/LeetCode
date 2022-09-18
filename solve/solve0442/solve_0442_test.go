package solve0442

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicates(t *testing.T) {
	assert.Equal(t, findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}), []int{2, 3})
	assert.Equal(t, findDuplicates([]int{1, 1, 2}), []int{1})
	assert.Equal(t, findDuplicates([]int{1}), []int{})
}
