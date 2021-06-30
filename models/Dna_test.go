package models

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDna(t *testing.T) {
	model := NewDna([]string{"1TGCGA", "*AGTGC", "-TATGT", "AGAAGG", "CCCCTA", "TCACTG",})
	require.IsType(t, &Dna{}, model)
}
