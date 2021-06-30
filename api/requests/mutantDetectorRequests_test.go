package requests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mutantDetector/models"
	"testing"
)

func TestNewMutantDetectorRequest(t *testing.T) {
	mdr := NewMutantDetectorRequest()
	require.IsType(t, &MutantDetectorRequest{}, mdr)
}

func TestValidate(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			true,
		},
		{
			[]string{"ATGCGA",},
			true,
		},
		{
			[]string{},
			false,
		},
	}

	for _, test := range tests {
		mdr := NewMutantDetectorRequest()
		mdr.Dna = test.input
		err := mdr.Validate()
		if test.expected {
			assert.NoError(t, err)
		} else {
			t.Log(err)
			assert.Error(t, err)
		}
	}
}

func TestToModel(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			true,
		},
		{
			[]string{},
			false,
		},
	}

	for _, test := range tests {
		mdr := NewMutantDetectorRequest()
		mdr.Dna = test.input
		modelDna := mdr.ToModel()
		assert.IsType(t, &models.Dna{}, modelDna )
	}
}
