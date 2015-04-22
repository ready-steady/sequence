package sampling

import (
	"errors"
)

// void sobol(int dim_num, long long int *seed, double quasi[]);
import "C"

// Sobol represents a Sobol sequence.
//
// https://en.wikipedia.org/wiki/Sobol_sequence
type Sobol struct {
	dims C.int
	seed C.longlong
}

// NewSobol returns a new Sobol sequence.
func NewSobol(dimensions uint, seed int64) (*Sobol, error) {
	const (
		DIM_MAX2 = 1111 // sobol/sobol.cpp:14280
	)

	if dimensions == 0 || dimensions > DIM_MAX2 {
		return nil, errors.New("the number of dimensions is invalid")
	}

	sobol := &Sobol{
		dims: C.int(dimensions),
		seed: C.longlong(seed),
	}

	return sobol, nil
}

// Next returns the next `count` elements of the sequence.
func (s *Sobol) Next(count uint) []float64 {
	dims, seed := s.dims, s.seed

	points := make([]float64, count*uint(dims))
	for i := uint(0); i < count; i++ {
		C.sobol(dims, &seed, (*C.double)(&points[i*uint(dims)]))
	}

	s.seed = seed

	return points
}
