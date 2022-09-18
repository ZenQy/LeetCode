package solve0645

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindErrorNums(t *testing.T) {
	assert.Equal(t, findErrorNums([]int{1,2,2,4}), []int{2,3})
	assert.Equal(t, findErrorNums([]int{1,1}),[]int{1,2})
}

