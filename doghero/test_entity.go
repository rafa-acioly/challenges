package doghero

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DogWalkingTestSuit struct {
	suite.Suite
}

func (suite *DogWalkingTestSuit) TestShow() {
	walking := NewWalk()

	walking.EndAt = time.Now().Add(time.Hour * 2)

	expected := 60

	assert.Equal(suite.T(), expected, walking.Show())
}

func TestDogWalkingEntity(t *testing.T) {
	suite.Run(t, new(DogWalkingTestSuit))
}
