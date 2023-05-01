package number

import "labTest/internal/ports"

type repo struct {
}

func New() ports.NumberRepository {
	return &repo{}
}

func (r *repo) GetN1() int {
	return 3
}

func (r *repo) GetN2() int {
	return 2
}
