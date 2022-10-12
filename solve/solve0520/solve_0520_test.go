package solve0520

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCapitalUse(t *testing.T) {
	assert.Equal(t, detectCapitalUse("USA"), true)
	assert.Equal(t, detectCapitalUse("word"), true)
	assert.Equal(t, detectCapitalUse("True"), true)
	assert.Equal(t, detectCapitalUse("FlaG"), false)
}
