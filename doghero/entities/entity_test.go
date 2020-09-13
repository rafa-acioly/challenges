package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DogWalkingTestSuit struct {
	suite.Suite
}

func (suite *DogWalkingTestSuit) TestNewWalk() {
	assert.Equal(suite.T(), DogWalking{}, NewWalk())
}

func (suite *DogWalkingTestSuit) TestShow() {
	walking := NewWalk()

	walking.StartAt = time.Now()
	walking.EndAt = time.Now().Add(time.Hour * 2)

	expected := 120.00

	assert.GreaterOrEqual(suite.T(), walking.Show(), expected)
}

func TestDogWalkingEntity(t *testing.T) {
	suite.Run(t, new(DogWalkingTestSuit))
}
