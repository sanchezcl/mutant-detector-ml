package repositories

import "mutantDetector/models"

type DnaStatsRepositoryInterface interface {
	GetStats() (*models.Stats, error)
}
