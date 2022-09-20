package solve0119

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
	assert.Equal(t, getRow(3), []int{1, 3, 3, 1})
	assert.Equal(t, getRow(1), []int{1, 1})
}
