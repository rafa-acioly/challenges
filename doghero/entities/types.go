package entities

type WalkingStatus int

const (
	WalkingPending WalkingStatus = iota
	WalkingInProgress
	WalkingFinished
)
