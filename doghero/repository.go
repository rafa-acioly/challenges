package doghero

import (
	"database/sql"
	"errors"
	"time"
)

type Repository interface {
	Index(filter WalkingFilter, page int) []DogWalking
	Create(walking DogWalking) bool
	StartWalk(walkingUUID string) (time.Time, error)
	FinishWalk(walkingUUID string) (time.Time, error)
}

type repository struct {
	database *sql.DB
}

// NewRepository retrieve a repository instance connection
// to the walking database
func NewRepository(db *sql.DB) Repository {
	return &repository{database: db}
}

// Index retrieve a list of walking
func (r repository) Index(filter WalkingFilter, page int) []DogWalking {
	return []DogWalking{}
}

// Create insert a new walking record on the database
func (r repository) Create(walking DogWalking) bool {
	return true
}

// StartWalk set the "StartAt" key to the current time if it is not defined yet
func (r *repository) StartWalk(walkingUUID string) (time.Time, error) {
	started := time.Now()

	query := "UPDATE table SET start_at = $1 WHERE uuid = $2 AND start_at = null"
	stmt, err := r.database.Prepare(query)
	if err != nil {
		return time.Time{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(started, walkingUUID)
	if err != nil {
		return time.Time{}, err
	}

	if row, _ := result.RowsAffected(); row == 0 {
		return time.Time{}, errors.New("walking already started or does not exist")
	}

	return started, nil
}

// FinishWalk set the "EndAt" key to the current time
func (r *repository) FinishWalk(walkingUUID string) (time.Time, error) {
	return time.Time{}, nil
}
