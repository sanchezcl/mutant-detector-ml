package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mutantDetector/models"
	"testing"
)

type DbTestConn struct {
	GormDb *gorm.DB
	Mock   sqlmock.Sqlmock
}

func getTestDb() (*DbTestConn, error) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, err
	}

	dialect := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDb, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dtc := &DbTestConn{
		GormDb: gormDb,
		Mock:   mock,
	}

	return dtc, nil
}

func TestNewPgDnaRepository(t *testing.T) {
	conn, err := getTestDb()
	require.NoError(t, err)
	repo := NewPgGormDnaRepository(conn.GormDb)
	require.IsType(t, &PgGormDnaRepository{}, repo)
}

func TestPgDnaRepository_Create(t *testing.T) {
	//t.Skip("Skipped because a bug with gorm and go-sqlmock or i'm doing something wrong and i can't see it")

	dnaModel := models.NewDna([]string{"1TGCGA", "*AGTGC", "-TATGT", "AGAAGG", "CCCCTA", "TCACTG",})
	dnaModel.IsMutant = true

	testDb, err := getTestDb()
	require.NoError(t, err)

	const sqlInsert = `INSERT INTO "dnas" ("id","created_at","updated_at","deleted_at","hash","is_mutant") VALUES ($1,$2,$3,$4,$5,$6)`

	testDb.Mock.MatchExpectationsInOrder(false)
	testDb.Mock.ExpectBegin()
	testDb.Mock.ExpectQuery(sqlInsert).
		WithArgs(
			sqlmock.AnyArg(), //uuid
			sqlmock.AnyArg(), //created_at
			sqlmock.AnyArg(), //updated_at
			sqlmock.AnyArg(), //deleted_at,
			dnaModel.Hash,
			dnaModel.IsMutant,
		).
		WillReturnRows(sqlmock.NewRows([]string{"hash"}).AddRow(dnaModel.Hash))
	testDb.Mock.ExpectCommit()

	repo := NewPgGormDnaRepository(testDb.GormDb)
	err2 := testDb.Mock.ExpectationsWereMet()
	var resultDna *models.Dna
	resultDna, err = repo.Create(dnaModel)
	require.NoError(t, err2)

	assert.IsType(t, &models.Dna{}, resultDna)
	assert.NotEmpty(t, resultDna)
}

func TestPgDnaRepository_FindByHash(t *testing.T) {
	t.Skip("To implement")
}
