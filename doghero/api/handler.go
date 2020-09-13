package api

import "github.com/rafa-acioly/challenges/repositories"

type Resource struct {
	repository repositories.Repository
}

func NewDogWalkResource() Resource {
	return Resource{repository: repositories.NewDogWalkRepository()}
}
