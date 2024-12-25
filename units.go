package forzatelemetry

import "time"

// Main temperature unit of the protocol
type Fahrenheit float32

func (f Fahrenheit) Fahrenheit() float32 {
	return float32(f)
}

func (f Fahrenheit) Celsius() float32 {
	return float32(f-32) / 1.8
}

// Main distance unit of the protocol
type Meter float32

func (m Meter) Meter() float32 {
	return float32(m)
}

func (m Meter) CentiMeter() float32 {
	return float32(m * 100)
}

func (m Meter) KiloMeter() float32 {
	return float32(m / 1000)
}

func (m Meter) Mile() float32 {
	return float32(m / 1609.344)
}

func (m Meter) Inch() float32 {
	return float32(m * 39.3700787)
}

// Main speed/velocity unit of the protocol
type MeterPerSecond float32

func (mps MeterPerSecond) MeterPerSecond() float32 {
	return float32(mps)
}

func (mps MeterPerSecond) KilometerPerHour() float32 {
	return float32(mps * 3.6)
}

func (mps MeterPerSecond) MilePerHour() float32 {
	return float32(mps * 2.23693629)
}

// Main angle unit of the protocol
type Radian float32

func (r Radian) Radian() float32 {
	return float32(r)
}

func (r Radian) Degree() float32 {
	return float32(r * 57.2957795)
}

// Main power unit of the protocol
type Watt float32

func (w Watt) Watt() float32 {
	return float32(w)
}

func (w Watt) KiloWatt() float32 {
	return float32(w / 1000)
}

func (w Watt) Horsepower() float32 {
	return float32(w / 745.699872)
}

// Main torque unit of the protocol
type NewtonMeter float32

func (nm NewtonMeter) NewtonMeter() float32 {
	return float32(nm)
}

func (nm NewtonMeter) FootPound() float32 {
	return float32(nm / 1.3558179483)
}

// Main pressure unit of the protocol
type PoundPerSuqaredInch float32

func (psi PoundPerSuqaredInch) PoundPerSuqaredInch() float32 {
	return float32(psi)
}

func (psi PoundPerSuqaredInch) Bar() float32 {
	return float32(psi / 14.5037738)
}

// Main time unit of the protocol
type Second float32

func (s Second) Second() float32 {
	return float32(s)
}

func (s Second) Duration() time.Duration {
	return time.Duration(s * 1000000000)
}

// Main timestamp unit of the protocol
type MilliSecond uint32

func (ms MilliSecond) MilliSecond() uint32 {
	return uint32(ms)
}

func (ms MilliSecond) Duration() time.Duration {
	return time.Duration(ms * 1000000)
}
