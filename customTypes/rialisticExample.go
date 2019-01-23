package main

import (
	"fmt"
)

type Meters float64
type MillMeter float64
type NanoMeter float64

func (m MillMeter) toMeter() Meters {
	return Meters(m / 1000)
}

func (n NanoMeter) toMeter() Meters {
	return Meters(n / 1e+9)
}

func main() {
	mi := MillMeter(2000)
	fmt.Printf("Converting %.1f millimeters to %.1f merter\n", mi, mi.toMeter())

	nm := NanoMeter(200000000)
	fmt.Printf("Converting %.1f millimeters to %.1f merter\n", nm, nm.toMeter())
}
