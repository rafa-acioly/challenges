package doghero

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var walkingStub = DogWalking{
	Status:      0,
	ScheduledTo: time.Now(),
	Price:       0,
	Duration:    0,
	Lat:         001,
	Long:        002,
	Pets:        0,
	StartAt:     time.Time{},
	EndAt:       time.Time{},
}

type DogWalkingTestSuit struct {
	suite.Suite
}

func (suite *DogWalkingTestSuit) SetupSuite() {

}

func (suite *DogWalkingTestSuit) TearDownTest() {

}

func (suite *DogWalkingTestSuit) TestShow() {
	walking := walkingStub

	walking.EndAt = time.Now().Add(time.Hour * 2)

	expected := 60

	assert.Equal(suite.T(), expected, walking.Show())
}

func (suite *DogWalkingTestSuit) TestStartWalk() {
	walk := walkingStub

	walk.StartWalk()

	assert.False(suite.T(), walk.StartAt.IsZero())
}

func (suite *DogWalkingTestSuit) TestFinishWalk() {
	walk := walkingStub

	walk.FinishWalk()

	assert.False(suite.T(), walk.EndAt.IsZero())
}

func TestDogWalkingEntity(t *testing.T) {
	suite.Run(t, new(DogWalkingTestSuit))
}
