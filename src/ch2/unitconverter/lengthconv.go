// Package unitconverter with support for length, weight and temperature conversion.
package unitconverter

import "fmt"

type Foot float64
type Meter float64

func (ft Foot) String() string { return fmt.Sprintf("%g ft", ft) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }

// FToM converts Foot reading to Meter
func FToM(ft Foot) Meter { return Meter(ft / 3.281) }

// MToF converts Meter reading to Foot
func MToF(m Meter) Foot { return Foot(m * 3.281) }
