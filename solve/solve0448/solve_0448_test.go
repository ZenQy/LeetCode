package solve0448

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDisappearedNumbers(t *testing.T) {
	assert.Equal(t, findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}), []int{5, 6})
	assert.Equal(t, findDisappearedNumbers([]int{1, 1}), []int{2})
	assert.Equal(t, findDisappearedNumbers([]int{1, 2}), []int{})
}
