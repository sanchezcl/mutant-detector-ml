package repositories

import (
	"gorm.io/gorm"
	"mutantDetector/models"
)

type PgGormDnaStatsRepository struct {
	dbConn *gorm.DB
}

func NewPgGormStatsRepository(conn *gorm.DB) *PgGormDnaStatsRepository {
	return &PgGormDnaStatsRepository{
		dbConn: conn,
	}
}

func (r *PgGormDnaStatsRepository) GetStats() (*models.Stats, error){
	var mutants, humans int64
	res := r.dbConn.Model(&models.Dna{}).Where("is_mutant = ?", true).Count(&mutants)
	if res.Error != nil {
		return nil, res.Error
	}
	res = r.dbConn.Model(&models.Dna{}).Where("is_mutant = ?", false).Count(&humans)
	if res.Error != nil {
		return nil, res.Error
	}
	return models.NewStats(mutants, humans), nil
}
