package doghero

type Resource struct {
	repository Repository
}

func NewResource() Resource {
	return Resource{repository: NewRepository()}
}
