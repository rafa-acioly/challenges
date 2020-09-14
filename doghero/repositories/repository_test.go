package repositories

import (
	"database/sql"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuit struct {
	suite.Suite
	repository Repository
	database   *sql.DB
	mock       sqlmock.Sqlmock
}

func (suite *RepositoryTestSuit) SetupSuite() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	suite.database = db

	suite.repository = NewDogWalkRepository(db)
}

func (suite *RepositoryTestSuit) TestIndexFilterAll() {
	suite.mock.
		ExpectQuery("SELECT \\* FROM walks").
		WillReturnRows(mockedRows)

	result := suite.repository.Index(All, 0)

	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
	assert.NotEmpty(suite.T(), result)
	assert.Len(suite.T(), result, 2)
}

func (suite *RepositoryTestSuit) TestIndexFilterNext() {
	suite.T().Skip()

	suite.mock.
		ExpectQuery("SELECT \\* FROM walks").
		WillReturnRows(mockedRows)

	result := suite.repository.Index(Next, 0)

	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
	assert.NotEmpty(suite.T(), result)
	assert.Len(suite.T(), result, 1)
}

func (suite *RepositoryTestSuit) TestCreate() {
	suite.mock.
		ExpectPrepare("INSERT INTO walks").
		ExpectExec().
		WithArgs(
			walkingStub.ID, walkingStub.Status,
			walkingStub.ScheduledTo, walkingStub.Price,
			walkingStub.Duration, walkingStub.Lat,
			walkingStub.Long, walkingStub.Pets,
			walkingStub.StartAt, walkingStub.EndAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := suite.repository.Create(walkingStub)

	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), result)
}

func (suite *RepositoryTestSuit) TestStartWalk() {
	suite.mock.
		ExpectPrepare("UPDATE TABLE walks SET start_at = \\? WHERE id = \\?").
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	period, err := suite.repository.StartWalk(walkingStub.ID)

	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), period)
}

func (suite *RepositoryTestSuit) TestFinishWalk() {
	suite.mock.
		ExpectPrepare("UPDATE TABLE walks SET end_at = \\? WHERE id = \\?").
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	period, err := suite.repository.FinishWalk(walkingStub.ID)

	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), period)
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuit))
}
