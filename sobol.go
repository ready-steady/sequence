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
		cursor[i] = uint32(scramble)
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
		size = 1 << bits
	)

	dimensions, offset, cursor := s.dimensions, s.offset, s.cursor

	index := make([]uint, points)
	for i := uint(0); i < points; i++ {
		for j := offset + i; j&1 != 0; j >>= 1 {
			index[i]++
		}
	}

	data := make([]float64, points*dimensions)
	for i := uint(0); i < dimensions; i++ {
		for j := uint(0); j < points; j++ {
			data[j*dimensions+i] = float64(cursor[i]) / size
			cursor[i] ^= sobolData[i*bits+index[j]]
		}
	}

	s.offset += points

	return data
}
