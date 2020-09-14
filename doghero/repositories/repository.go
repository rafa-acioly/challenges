package repositories

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/rafa-acioly/challenges/entities"
)

// Repository represents a contract
type Repository interface {
	Index(filter WalkingFilter, page int) []entities.DogWalking
	Create(walking entities.DogWalking) (bool, error)
	StartWalk(walkingUUID string) (time.Time, error)
	FinishWalk(walkingUUID string) (time.Time, error)
}

type dogWalkRepository struct {
	database *sql.DB
}

// NewDogWalkRepository retrieve a repository instance connection
// to the walking database
func NewDogWalkRepository(db *sql.DB) Repository {
	return &dogWalkRepository{database: db}
}

// Index retrieve a list of walking
func (r dogWalkRepository) Index(filter WalkingFilter, page int) []entities.DogWalking {
	query := r.filterQuery(filter)
	rows, _ := r.database.Query(query)
	defer rows.Close()

	var results []entities.DogWalking
	for rows.Next() {
		var entity entities.DogWalking
		if err := rows.Scan(&entity.ID, &entity.Status, &entity.ScheduledTo); err != nil {
			log.Fatal(err.Error())
		}

		results = append(results, entity)
	}

	return results
}

func (r dogWalkRepository) filterQuery(filter WalkingFilter) string {
	if filter == 0 {
		return "SELECT * FROM walks LIMIT 20"
	}

	return "SELECT * FROM walks WHERE scheduled_to >= NOW() LIMIT 20"
}

// Create insert a new walking record on the database
func (r dogWalkRepository) Create(walking entities.DogWalking) (bool, error) {
	stmt, err := r.database.Prepare("INSERT INTO walks VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return false, err
	}

	result, err := stmt.Exec(
		walking.ID, walking.Status,
		walking.ScheduledTo, walking.Price,
		walking.Duration, walking.Lat,
		walking.Long, walking.Pets,
		walking.StartAt, walking.EndAt,
	)
	if err != nil {
		return false, err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return false, errors.New("line not inserted on the database")
	}

	return true, nil
}

// StartWalk set the "StartAt" key to the current time and retrieve the time defined
func (r *dogWalkRepository) StartWalk(walkingUUID string) (time.Time, error) {
	stmt, err := r.database.Prepare("UPDATE TABLE walks SET start_at = ? WHERE id = ?")
	if err != nil {
		return time.Time{}, err
	}

	currentTime := time.Now()
	if _, err := stmt.Exec(currentTime, walkingUUID); err != nil {
		return time.Time{}, err
	}

	return currentTime, nil
}

// FinishWalk set the "EndAt" key to the current time and retrieve the time defined
func (r *dogWalkRepository) FinishWalk(walkingUUID string) (time.Time, error) {
	stmt, err := r.database.Prepare("UPDATE TABLE walks SET end_at = ? WHERE id = ?")
	if err != nil {
		return time.Time{}, err
	}

	currentTime := time.Now()
	if _, err := stmt.Exec(walkingUUID, currentTime); err != nil {
		return time.Time{}, err
	}

	return currentTime, nil
}
