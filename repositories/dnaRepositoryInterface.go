package repositories

import "mutantDetector/models"

type DnaRepositoryInterface interface {
	FindByHash(dna *models.Dna) (*models.Dna, error)
	Create(dna *models.Dna) (*models.Dna, error)
}
