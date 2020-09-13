package repositories


type WalkingFilter int

// TODO(rafa): This should be a unique "Enum" to identify the filter types.
const (
	All WalkingFilter = iota
	Next
)
