package doghero

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

	suite.repository = NewRepository(db)
}

func (suite *RepositoryTestSuit) TestNewRepository() {
}

func (suite *RepositoryTestSuit) TestIndexFilterAll() {
	assert.Empty(suite.T(), suite.repository.Index(All, 0))
}

func (suite *RepositoryTestSuit) TestIndexFilterNext() {
	assert.Empty(suite.T(), suite.repository.Index(Next, 0))
}

func (suite *RepositoryTestSuit) TestCreate() {
	suite.mock.
		ExpectExec("INSERT INTO x").
		WillReturnResult(sqlmock.NewResult(1, 1))

	created := suite.repository.Create(DogWalking{})

	assert.True(suite.T(), created)
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *RepositoryTestSuit) TestStartWalk() {
	suite.mock.
		ExpectExec("UPDATE table").
		WithArgs("date", "uuid").
		WillReturnResult(sqlmock.NewResult(1, 1))

	period, err := suite.repository.StartWalk("uuid")

	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), period)
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *RepositoryTestSuit) TestCannotStartWalkAlreadyStarted() {
	suite.mock.
		ExpectExec("UPDATE table").
		WithArgs("date", "uuid").
		WillReturnResult(sqlmock.NewResult(0, 0))

	period, err := suite.repository.StartWalk("uuid")

	assert.Error(suite.T(), err)
	assert.Zero(suite.T(), period)
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *RepositoryTestSuit) TestFinishWalk() {

}

func (suite *RepositoryTestSuit) TestCannotFinishWithoutStarting() {

}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuit))
}
