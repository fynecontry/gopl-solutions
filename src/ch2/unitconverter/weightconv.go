// Package unitconverter with support for length, weight and temperature conversion.
package unitconverter

import "fmt"

type Pound float64
type Kilogram float64

func (lb Pound) String() string    { return fmt.Sprintf("%g lb", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%g kg", kg) }

// LBtoK converts pound to kilogram.
func LBtoKg(lb Pound) Kilogram { return Kilogram(lb / 2.205) }

// KgtoLB converts Kilogram to Pound.
func KgtoLB(kg Kilogram) Pound { return Pound(kg * 2.205) }
