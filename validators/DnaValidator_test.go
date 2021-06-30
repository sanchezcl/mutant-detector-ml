package validators

import (
	"github.com/stretchr/testify/require"
	"mutantDetector/models"
	"testing"
)

func TestNewMutantDetectorService(t *testing.T) {
	v := NewDnaValidator(&models.Dna{})
	require.IsType(t, &DnaValidator{}, v)
}

func TestValidate(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			true,
		},
		{
			[]string{"1TGCGA", "*AGTGC", "-TATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			false,
		},
		{
			[]string{"ATGCGA", "CAGGC", "TTATTT", "AGACGG", "GCGTCA", "TCCCACTG"},
			false,
		},
		{
			[]string{},
			false,
		},
	}

	for _, test := range tests {
		model := &models.Dna{
			M: test.input.([]string),
		}
		validator := NewDnaValidator(model)
		result, err := validator.Validate()
		if !test.expected {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, test.expected, result)
		}
	}
}

func TestValidateEmpty(t *testing.T) {
	tests := []struct {
		input    interface{}
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
		model := &models.Dna{
			M: test.input.([]string),
		}
		validator := NewDnaValidator(model)
		result, err := validator.validateEmpty()
		if !test.expected {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, test.expected, result)

	}
}

func TestValidateLetters(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			true,
		},
		{
			[]string{"ATGCGA", "CAGRGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"},
			false,
		},
		{
			[]string{"1TGCGA", "*AGTGC", "-TATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			false,
		},
	}

	for _, test := range tests {
		model := &models.Dna{
			M: test.input.([]string),
		}
		validator := NewDnaValidator(model)
		result, err := validator.validateLetters()
		if err != nil {
			require.Error(t, err, "incorrect letters")
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, test.expected, result)
	}
}

func TestValidateRowLen(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			true,
		},
		{
			[]string{"ATGCGA", "CAGGC", "TTATTT", "AGACGG", "GCGTCA", "TCCCACTG"},
			false,
		},
	}

	for _, test := range tests {
		model := &models.Dna{
			M: test.input.([]string),
		}
		validator := NewDnaValidator(model)
		result, err := validator.validateRowLen()
		if err != nil {
			require.Error(t, err, "wrong matrix length")
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, test.expected, result)
	}
}