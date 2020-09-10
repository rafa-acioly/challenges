package doghero

type WalkingStatus int

const (
	WalkingPending WalkingStatus = iota
	WalkingInProgress
	WalkingFinished
)

type WalkingFilter int

const (
	All WalkingFilter = iota
	Next
)
