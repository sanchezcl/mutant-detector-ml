package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStats(t *testing.T) {
	tests := []struct {
		m int64
		h int64
		r float32
	}{{3, 4, float32(3)/float32(4)}, {0, 4, float32(0)}, {4, 0, float32(4.0)},}

	for _, test := range tests {
		model := NewStats(test.m, test.h)
		require.IsType(t, &Stats{}, model)
		assert.Equal(t, test.m, model.CountMutantDna)
		assert.Equal(t, test.h, model.CountHumanDna)
		assert.Equal(t, test.r, model.Ratio)
	}
}

