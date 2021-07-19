package database

import (
	"gorm.io/gorm"
	"mutantDetector/config"
	"mutantDetector/models"
	"sync"
)

var (
	instance *gorm.DB
	onceDb   sync.Once
)

func NewDatabaseConn() *gorm.DB {
	onceDb.Do(func() {
		instance = getInstance()
	})
	return instance
}

func getInstance() *gorm.DB {
	pgDb := GetPgConn(getDbConfig())
	return pgDb
}

func getDbConfig() *config.Database {
	c := config.NewConfig()
	return &c.Db
}

func Migrate() {
	db := NewDatabaseConn()
	err := db.AutoMigrate(&models.Dna{})
	if err != nil {
		panic("couldn't run migrations")
	}
}
