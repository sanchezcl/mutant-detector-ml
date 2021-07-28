package models

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDna(t *testing.T) {
	model := NewDna([]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",})
	require.IsType(t, &Dna{}, model)
}

func TestParseDnaHash(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			"6d32697470c08f971ea2f5a71113166f9abc2d7f",
		},
		{
			[]string{"ATGCGA", "CAGAGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"},
			"91cf2394a355f8aa5e696ef3c3a2311ca7ed5414",
		},
	}

	for _, test := range tests {
		dna := NewDna(test.input)
		hash := dna.parseDnaHash()
		require.Equal(t, test.expected, hash)
	}
}