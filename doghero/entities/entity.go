package entities

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/is"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/google/uuid"
)

type DogWalking struct {
	ID          string        `json:"id"`
	Status      WalkingStatus `json:"status"`
	ScheduledTo time.Time     `json:"scheduled_to"`
	Price       float64       `json:"price"`
	Duration    int32         `json:"duration"`
	Lat         string        `json:"lat"`
	Long        string        `json:"long"`
	Pets        int           `json:"pets"`
	StartAt     time.Time     `json:"start_at"`
	EndAt       time.Time     `json:"end_at"`
}

// NewWalk retrieves a new walking instance with
// some default values pre-defined
func NewWalk() DogWalking {
	return DogWalking{
		ID:     uuid.New().String(),
		Status: WalkingPending,
	}
}

// Show retrieves the walk duration in minutes
func (d *DogWalking) Show() float64 {
	return d.EndAt.Sub(d.StartAt).Minutes()
}

// Valid validates the required rules for a walking
func (d DogWalking) Valid() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.ScheduledTo, validation.Required, validation.Min(time.Now())),
		validation.Field(&d.Lat, validation.Required, is.Latitude),
		validation.Field(&d.Long, validation.Required, is.Longitude),
		validation.Field(&d.Pets, validation.Required, validation.Min(1)),
	)
}
