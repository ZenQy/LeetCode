package solve0451

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	assert.Equal(t, "eert", frequencySort("tree"))
	assert.Equal(t, frequencySort("cccaaa"), "cccaaa")
	assert.Equal(t, frequencySort("Aabb"), "bbAa")
}
