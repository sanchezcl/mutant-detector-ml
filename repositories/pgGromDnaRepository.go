package repositories

import (
	"errors"
	"gorm.io/gorm"
	"mutantDetector/models"
)

type PgGormDnaRepository struct{
	dbConn *gorm.DB
}

func NewPgGormDnaRepository(conn *gorm.DB) *PgGormDnaRepository {
	return &PgGormDnaRepository{
		dbConn: conn,
	}
}

func (r *PgGormDnaRepository) FindByHash(dna *models.Dna) (*models.Dna, error) {
	var resDna models.Dna
	res := r.dbConn.Limit(1).Find(&resDna, "hash = ?", dna.Hash)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &resDna, nil
}

func (r *PgGormDnaRepository) Create(dna *models.Dna) (*models.Dna, error) {
	res := r.dbConn.Save(dna)
	if res.Error != nil {
		return nil, res.Error
	}
	return dna, nil
}
