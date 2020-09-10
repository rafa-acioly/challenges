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
	suite.repository = NewRepository()
}

func (suite *RepositoryTestSuit) TestNewRepository() {
	assert.Equal(suite.T(), Repository{}, NewRepository())
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

func (suite *RepositoryTestSuit) TestFinishWalk() {

}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuit))
}
