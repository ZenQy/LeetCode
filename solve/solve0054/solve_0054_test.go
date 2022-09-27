package solve0054

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	assert.Equal(t, spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}}), []int{1,2,3,6,9,8,7,4,5})
	assert.Equal(t, spiralOrder([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}), []int{1,2,3,4,8,12,11,10,9,5,6,7})
}

