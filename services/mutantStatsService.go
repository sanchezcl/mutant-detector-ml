package services

import (
	"mutantDetector/models"
	"mutantDetector/repositories"
)

type MutantStatsService struct {
	repo repositories.DnaStatsRepositoryInterface
}

func NewMutantStatsService(repo repositories.DnaStatsRepositoryInterface) (*MutantStatsService, error){
	return &MutantStatsService{
		repo:  repo,
	}, nil
}

func (mss *MutantStatsService) GetStats() (*models.Stats, error){
	stats, err := mss.repo.GetStats()
	if err != nil {
		return nil, err
	}
	return stats, nil
}
