package repositories

import (
	"time"

	"github.com/rafa-acioly/challenges/entities"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var walkingStub = entities.DogWalking{
	ID:          "walk-id",
	Status:      entities.WalkingPending,
	ScheduledTo: time.Now().AddDate(0, 0, 1), // tomorrow
	Price:       25,
	Duration:    30,
	Lat:         0,
	Long:        0,
	Pets:        1,
	StartAt:     time.Time{},
	EndAt:       time.Time{},
}

// walkingInPast represents a walking that is already finished
var walkingInPast = entities.DogWalking{
	ID:          "random-id",
	Status:      entities.WalkingFinished,
	ScheduledTo: time.Now().AddDate(0, 0, -2), // 2 days ago
	Price:       25,
	Duration:    30,
	Lat:         0,
	Long:        0,
	Pets:        1,
	StartAt:     time.Now().AddDate(0, 0, -1), // yesterday
	EndAt:       time.Now().AddDate(0, 0, -1), // yesterday
}

// walkInFuture represents a walking that is scheduled and pending
var walkInFuture = entities.DogWalking{
	ID:          "random-future-id",
	Status:      entities.WalkingPending,
	ScheduledTo: time.Now().AddDate(0, 0, 2), // 2 days
	Price:       25,
	Duration:    30,
	Lat:         0,
	Long:        0,
	Pets:        1,
	StartAt:     time.Time{},
	EndAt:       time.Time{},
}

// mockedRows is a mocked database records, it has just a few fields to keep
// this file clean
var mockedRows = sqlmock.NewRows([]string{"id", "status", "scheduled_to"}).
	AddRow(walkingInPast.ID, walkingInPast.Status, walkingInPast.ScheduledTo).
	AddRow(walkInFuture.ID, walkInFuture.Status, walkInFuture.ScheduledTo)
