package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mutantDetector/config"
)


func GetPgConn(c *config.Database) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(getDsn(c)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return conn
}

func getDsn(c *config.Database) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		c.DbHost,
		c.DbUsername,
		c.DbPassword,
		c.DbDatabase,
		c.DbPort)
}