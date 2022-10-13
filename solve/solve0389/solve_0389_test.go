package solve0389

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheDifference(t *testing.T) {
	assert.Equal(t, findTheDifference("abcd", "abcde"), byte('e'))
	assert.Equal(t, findTheDifference("", "y"), byte('y'))
}
