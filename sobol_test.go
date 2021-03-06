package sequence

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestNewSobol(t *testing.T) {
	sequence := NewSobol(3, 0x0102030405060708)
	assert.Equal(sequence.cursor, []uint32{0x05060708, 0x01020304, 0x05060708}, t)
}

func TestSobol3D(t *testing.T) {
	sequence := NewSobol(3, 0)

	actual := append(append([]float64{},
		sequence.Next(5)...),
		sequence.Next(5)...)

	expected := []float64{
		0.0000, 0.0000, 0.0000,
		0.5000, 0.5000, 0.5000,
		0.7500, 0.2500, 0.2500,
		0.2500, 0.7500, 0.7500,
		0.3750, 0.3750, 0.6250,
		0.8750, 0.8750, 0.1250,
		0.6250, 0.1250, 0.8750,
		0.1250, 0.6250, 0.3750,
		0.1875, 0.3125, 0.9375,
		0.6875, 0.8125, 0.4375,
	}

	assert.Equal(actual, expected, t)
}

func TestSobol10D(t *testing.T) {
	sequence := NewSobol(10, 0)

	actual := append(append(append(append([]float64{},
		sequence.Next(1)...),
		sequence.Next(2)...),
		sequence.Next(3)...),
		sequence.Next(4)...)

	expected := []float64{
		0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
		0.5000, 0.5000, 0.5000, 0.5000, 0.5000, 0.5000, 0.5000, 0.5000, 0.5000, 0.5000,
		0.7500, 0.2500, 0.2500, 0.2500, 0.7500, 0.7500, 0.2500, 0.7500, 0.7500, 0.7500,
		0.2500, 0.7500, 0.7500, 0.7500, 0.2500, 0.2500, 0.7500, 0.2500, 0.2500, 0.2500,
		0.3750, 0.3750, 0.6250, 0.8750, 0.3750, 0.1250, 0.3750, 0.8750, 0.8750, 0.6250,
		0.8750, 0.8750, 0.1250, 0.3750, 0.8750, 0.6250, 0.8750, 0.3750, 0.3750, 0.1250,
		0.6250, 0.1250, 0.8750, 0.6250, 0.6250, 0.8750, 0.1250, 0.1250, 0.1250, 0.3750,
		0.1250, 0.6250, 0.3750, 0.1250, 0.1250, 0.3750, 0.6250, 0.6250, 0.6250, 0.8750,
		0.1875, 0.3125, 0.9375, 0.4375, 0.5625, 0.3125, 0.4375, 0.9375, 0.9375, 0.3125,
		0.6875, 0.8125, 0.4375, 0.9375, 0.0625, 0.8125, 0.9375, 0.4375, 0.4375, 0.8125,
	}

	assert.Equal(actual, expected, t)
}
