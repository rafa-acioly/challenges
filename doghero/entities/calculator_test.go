package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CalculatorTestSuite struct {
	suite.Suite
}

func (suite *CalculatorTestSuite) TestCalculateOnePetThirtyMinutes() {
	pets, minutes := 1, 30.0

	price := GetWalkPrice(pets, minutes) // 1 pet and 30 minutes

	expectedPrice := 25

	assert.Equal(suite.T(), expectedPrice, price)
}

func (suite *CalculatorTestSuite) TestCalculateNPetsThirtyMinutes() {
	pets, minutes := 4, 30.0

	price := GetWalkPrice(pets, minutes)

	expectedPrice := 70

	assert.Equal(suite.T(), expectedPrice, price)
}

func (suite *CalculatorTestSuite) TestCalculateOnePetSixtyMinutes() {
	pets, minutes := 1, 60.0

	price := GetWalkPrice(pets, minutes)

	expectedPrice := 30

	assert.Equal(suite.T(), expectedPrice, price)
}

func (suite *CalculatorTestSuite) TestCalculateNPetsSixtyMinutes() {
	pets, minutes := 4, 60.0

	price := GetWalkPrice(pets, minutes)

	expectedPrice := 90

	assert.Equal(suite.T(), expectedPrice, price)
}

func TestCalculator(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
