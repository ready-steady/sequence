package sequence

// Sobol generates a Sobol sequence. The maximal number of dimensions is 21201,
// and the maximal number of points is 2^32.
//
// https://en.wikipedia.org/wiki/Sobol_sequence
func Sobol(dimensions, points uint, seed int64) []float64 {
	data := make([]float64, points*dimensions)

	index := make([]uint, points)
	for i := uint(1); i < points; i++ {
		for j := i - 1; j&1 != 0; j >>= 1 {
			index[i]++
		}
	}

	for i := uint(0); i < dimensions; i++ {
		data[i] = float64(uint32(seed)) / (1 << sobolBits)
		for j, x := uint(1), uint32(seed); j < points; j++ {
			x ^= sobolData[i*sobolBits+index[j]]
			data[j*dimensions+i] = float64(x) / (1 << sobolBits)
		}
	}

	return data
}
