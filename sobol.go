package sequence

// Sobol represents a Sobol sequence. The maximal number of dimensions is 21201,
// and the maximal number of points is 2^32.
//
// https://en.wikipedia.org/wiki/Sobol_sequence
type Sobol struct {
	dimensions uint
	offset     uint
	cursor     []uint32
}

// NewSobol returns a new Sobol sequence.
func NewSobol(dimensions uint, scramble int64) *Sobol {
	cursor := make([]uint32, dimensions)
	for i := range cursor {
		if i%2 == 0 {
			cursor[i] = uint32(scramble)
		} else {
			cursor[i] = uint32(scramble >> 32)
		}
	}

	return &Sobol{
		dimensions: dimensions,
		cursor:     cursor,
	}
}

// Next advances the sequence and returns the traversed points.
func (s *Sobol) Next(points uint) []float64 {
	const (
		bits = 32
	)

	dimensions, offset, cursor := s.dimensions, s.offset, s.cursor

	data := make([]float64, points*dimensions)
	for i := uint(0); i < points; i++ {
		k := uint(0)
		for j := offset + i; j&1 != 0; j >>= 1 {
			k++
		}
		for j := uint(0); j < dimensions; j++ {
			data[i*dimensions+j] = float64(cursor[j]) / (1 << bits)
			cursor[j] ^= sobolData[j*bits+k]
		}
	}

	s.offset += points

	return data
}
