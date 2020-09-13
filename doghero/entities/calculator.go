package entities

import (
	"time"
)

const (
	// ThirtyMinutesPrice 30 minutes & 1 pet
	ThirtyMinutesPrice = 25 // R$25.00
	// ThirtyMinutesAditionalPrice aditional value from more then 1 pet
	ThirtyMinutesAditionalPrice = 15 // R$5.00

	// SixtyMinutesPrice 60 minutes & 1 pet
	SixtyMinutesPrice = 30 // R$35.00
	// SixtyMinutesAditionalPrice adicional value for more then 1 pet
	SixtyMinutesAditionalPrice = 20 // R$10.00
)

// GetWalkPrice calculates the walking price
func GetWalkPrice(petsQuantity int, duration float64) int {
	thirtyMinutes := time.Duration(time.Minute * 30).Minutes()

	calculation := calculateThirtyMinutes
	if duration > thirtyMinutes {
		calculation = calculateSixtyMinutes
	}

	return calculation(petsQuantity)
}

// calculateThirtyMinutes will calculate the walking price based
// on the thirty minutes walk price
func calculateThirtyMinutes(petsQuantity int) int {
	extraPets := (petsQuantity - 1)

	return ThirtyMinutesPrice + (extraPets * ThirtyMinutesAditionalPrice)
}

// calculateSixtyMinute will calculate the walking price based
// on the sixty minutes walk price
func calculateSixtyMinutes(petsQuantity int) int {
	extraPets := (petsQuantity - 1)

	return SixtyMinutesPrice + (extraPets * SixtyMinutesAditionalPrice)
}
