package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVerseIDCumSum(t *testing.T) {
	result := getVerseIDCumSum()
	assert.Len(t, result, nChapters)
	assert.Equal(t, 7, result[0])
	assert.Equal(t, nVerses, result[nChapters-1])
	assert.IsIncreasing(t, result)
}
