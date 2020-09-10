package doghero

import "time"

const collectionName = "dog_walking"

type Repository struct{}

// NewRepository retrieve a repository instance connection
// to the walking database
func NewRepository() Repository {
	return Repository{}
}

// Index retrieve a list of walking
func (r *Repository) Index(filter WalkingFilter, page int) []DogWalking {
	return []DogWalking{}
}

// Create insert a new walking record on the database
func (r *Repository) Create(walking DogWalking) bool {
	return true
}

// StartWalk set the "StartAt" key to the current time
func (r *Repository) StartWalk(walkingID int) (time.Time, error) {
	return time.Time{}, nil
}

// FinishWalk set the "EndAt" key to the current time
func (r *Repository) FinishWalk(walkingID int) (time.Time, error) {
	return time.Time{}, nil
}
