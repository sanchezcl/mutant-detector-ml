package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mutantDetector/models"
	"testing"
)

func TestNewMutantDetectorService(t *testing.T) {
	dna := models.NewDna([]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",})
	service, err := NewMutantDetectorService(dna)
	assert.NoError(t, err)
	assert.IsType(t, &MutantDetectorService{}, service)
}

func TestGetVerticalsExtraction(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			[]string{"ACTACT", "TATGCC", "GGAACA", "CTTACC", "GGGGTT", "ACTGAG"},
		},
		{
			[]string{"ATGCGA", "CAGAGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"},
			[]string{"ACTAGT", "TATGCC", "GGAAGA", "CATCTC", "GGTGCT", "ACTGAG"},
		},
	}

	for _, test := range tests {
		dna := models.NewDna(test.input.([]string))
		mds, err := NewMutantDetectorService(dna)
		require.NoError(t, err)
		result, err := mds.getVerticals()
		require.NoError(t, err)
		require.Equal(t, test.expected, result)
	}
}

func TestGetHorizontalsExtraction(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
		},
	}

	for _, test := range tests {
		dna := models.NewDna(test.input.([]string))
		mds, err := NewMutantDetectorService(dna)
		require.NoError(t, err)
		result, err := mds.getHorizontals()
		require.NoError(t, err)
		require.Equal(t, test.expected, result)
	}
}

func TestGetDiagonalsExtraction(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			[]string{"AAAATG", "AGTACT"},
		},
		{
			[]string{"ATGCG", "CAGTG", "TTATG", "AGAAG", "CCCCT",},
			[]string{"AAAAT", "GTAGC"},
		},
	}

	for _, test := range tests {
		dna := models.NewDna(test.input.([]string))
		mds, _ := NewMutantDetectorService(dna)
		d, err := mds.getDiagonals()
		require.NoError(t, err)
		require.Equal(t, test.expected.([]string)[0], d[0])
		require.Equal(t, test.expected.([]string)[1], d[1])
	}
}

func TestDnaAnalyzer(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},
			true,
		},
		{
			[]string{"CTGCGA", "CAGTCT", "CTACAT", "CGCCGG", "TCCCAA", "CCACTA",},
			true,
		},
		{
			[]string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG",},
			false,
		},
	}
	for _, test := range tests {
		dna := models.NewDna(test.input)
		mds, err := NewMutantDetectorService(dna)
		require.NoError(t, err)
		result, err := mds.AnalyzeDna()
		require.NoError(t, err)
		require.Equal(t, test.expected, result)
	}
}

func BenchmarkDnaAnalyzer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dna := models.NewDna([]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG",},)
		_, _ = NewMutantDetectorService(dna)
	}
}