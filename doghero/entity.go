package doghero

import (
	"time"
)

type WalkingStatus int

const (
	WalkingPending WalkingStatus = iota
	WalkingInProgress
	WalkingFinished
)

type DogWalking struct {
	Status      WalkingStatus `json:"status"`
	ScheduledTo time.Time     `json:"scheduled_to"`
	Price       float64       `json:"price"`
	Duration    int32         `json:"duration"`
	Lat         int           `json:"lat"`
	Long        int           `json:"long"`
	Pets        int           `json:"pets"`
	StartAt     time.Time     `json:"start_at"`
	EndAt       time.Time     `json:"end_at"`
}

// NewWalk retrieves a new walking instance with
// some default values setted
func NewWalk() *DogWalking {
	return &DogWalking{Status: WalkingPending}
}

// Show retrieves the walk duration in minutes
func (d *DogWalking) Show() float64 {
	return d.StartAt.Sub(d.EndAt).Minutes()

}
