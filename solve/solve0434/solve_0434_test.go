package solve0434

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSegments(t *testing.T) {
	assert.Equal(t, countSegments("Hello, my name is John"), 5)
	assert.Equal(t, countSegments("Hello, my   name is John"), 5)
	assert.Equal(t, countSegments("  Hello,  my name is John  "), 5)
}
