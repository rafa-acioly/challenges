package doghero

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuit struct {
	suite.Suite
	repository Repository
}

func (suite *RepositoryTestSuit) SetupSuite() {
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

}

func (suite *RepositoryTestSuit) TestStartWalk() {

}

func (suite *RepositoryTestSuit) TestCannotStartWalkAlreadyStarted() {

}

func (suite *RepositoryTestSuit) TestFinishWalk() {

}

func (suite *RepositoryTestSuit) TestCannotFinishWithoutStarting() {

}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuit))
}
