package services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mutantDetector/models"
	"mutantDetector/repositories"
	"testing"
)

func TestNewMutantStatsService(t *testing.T) {
	mCtrl := gomock.NewController(t)
	defer mCtrl.Finish()
	repo := repositories.NewMockDnaStatsRepositoryInterface(mCtrl)

	service, err := NewMutantStatsService(repo)
	require.NoError(t, err)
	assert.IsType(t, &MutantStatsService{}, service)
}

func TestGetStats(t *testing.T) {
	tests := []models.Stats {
		models.Stats{
			CountMutantDna: 1,
			CountHumanDna:  1,
			Ratio:          1,
		},
		models.Stats{
			CountMutantDna: 2,
			CountHumanDna:  1,
			Ratio:          2,
		},
		models.Stats{
			CountMutantDna: 1,
			CountHumanDna:  2,
			Ratio:          0.5,
		},
	}
	mCtrl := gomock.NewController(t)
	defer mCtrl.Finish()
	repo := repositories.NewMockDnaStatsRepositoryInterface(mCtrl)

	for _, test := range tests {
		repo.EXPECT().GetStats().Return(&models.Stats{
			CountMutantDna: test.CountMutantDna,
			CountHumanDna:  test.CountHumanDna,
			Ratio:          test.Ratio,
		}, nil)
		service, err := NewMutantStatsService(repo)
		require.NoError(t, err)

		result, err := service.GetStats()
		require.NoError(t, err)
		assert.IsType(t, &models.Stats{}, result)
		assert.Equal(t, result.CountMutantDna, test.CountMutantDna)
		assert.Equal(t, result.CountHumanDna, test.CountHumanDna)
		assert.Equal(t, result.Ratio, test.Ratio)
	}
}
