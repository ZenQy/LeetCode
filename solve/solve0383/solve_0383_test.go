package solve0383

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	assert.Equal(t, canConstruct("a", "b"), false)
	assert.Equal(t, canConstruct("aa", "ab"), false)
	assert.Equal(t, canConstruct("aa", "aab"), true)
}
