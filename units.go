package forzatelemetry

import "time"

// Main temperature unit of the protocol
type Fahrenheit float32

// Gets the temperature in Fahrenheit.
func (f Fahrenheit) Fahrenheit() float32 {
	return float32(f)
}

// Gets the temperature in Celsius.
func (f Fahrenheit) Celsius() float32 {
	return float32((f - 32) / 1.8)
}

// Creates a temperature from Celsius.
func Celsius(c float32) Fahrenheit {
	return Fahrenheit(c*1.8 + 32)
}

// Main distance unit of the protocol
type Meter float32

// Gets the distance in meters.
func (m Meter) Meter() float32 {
	return float32(m)
}

// Gets the distance in centimeters.
func (m Meter) CentiMeter() float32 {
	return float32(m * 100)
}

// Creates a distance from centimeters.
func CentiMeter(cm float32) Meter {
	return Meter(cm / 100)
}

// Gets the distance in kilometers.
func (m Meter) KiloMeter() float32 {
	return float32(m / 1000)
}

// Creates a distance from kilometers.
func KiloMeter(km float32) Meter {
	return Meter(km * 1000)
}

// Gets the distance in miles.
func (m Meter) Mile() float32 {
	return float32(m / 1609.344)
}

// Creates a distance from miles.
func Mile(m float32) Meter {
	return Meter(m * 1609.344)
}

// Gets the distance in inches.
func (m Meter) Inch() float32 {
	return float32(m * 39.3700787)
}

// Creates a distance from inches.
func Inch(i float32) Meter {
	return Meter(i / 39.3700787)
}

// Main speed/velocity unit of the protocol
type MeterPerSecond float32

// Gets the speed in meters per second.
func (mps MeterPerSecond) MeterPerSecond() float32 {
	return float32(mps)
}

// Gets the speed in kilometers per hour.
func (mps MeterPerSecond) KilometerPerHour() float32 {
	return float32(mps * 3.6)
}

// Creates a speed from kilometers per hour.
func KilometerPerHour(kmph float32) MeterPerSecond {
	return MeterPerSecond(kmph / 3.6)
}

// Gets the speed in miles per hour.
func (mps MeterPerSecond) MilePerHour() float32 {
	return float32(mps * 2.23693629)
}

// Creates a speed from miles per hour.
func MilePerHour(mph float32) MeterPerSecond {
	return MeterPerSecond(mph / 2.23693629)
}

// Main angle unit of the protocol
type Radian float32

// Gets the angle in Radian.
func (r Radian) Radian() float32 {
	return float32(r)
}

// Gets the angle in degrees.
func (r Radian) Degree() float32 {
	return float32(r * 57.2957795)
}

// Creates an angle from degrees.
func Degree(d float32) Radian {
	return Radian(d / 57.2957795)
}

// Main power unit of the protocol
type Watt float32

// Gets the power in watts.
func (w Watt) Watt() float32 {
	return float32(w)
}

// Gets the power in kilowatts.
func (w Watt) KiloWatt() float32 {
	return float32(w / 1000)
}

// Creates a power from kilowatts.
func KiloWatt(kw float32) Watt {
	return Watt(kw * 1000)
}

// Gets the power in horsepower.
func (w Watt) Horsepower() float32 {
	return float32(w / 745.699872)
}

// Creates a power from horsepower.
func Horsepower(kw float32) Watt {
	return Watt(kw * 745.699872)
}

// Main torque unit of the protocol
type NewtonMeter float32

// Gets the torque in Newton meters.
func (nm NewtonMeter) NewtonMeter() float32 {
	return float32(nm)
}

// Gets the torque in pound feet.
func (nm NewtonMeter) PoundFoot() float32 {
	return float32(nm / 1.3558179483)
}

// Creates a power from pound feet.
func PoundFoot(lbft float32) NewtonMeter {
	return NewtonMeter(lbft * 1.3558179483)
}

// Main pressure unit of the protocol
type PoundPerSuqaredInch float32

// Gets the pressure in pound per inch squared.
func (psi PoundPerSuqaredInch) PoundPerSuqaredInch() float32 {
	return float32(psi)
}

// Gets the pressure in bars.
func (psi PoundPerSuqaredInch) Bar() float32 {
	return float32(psi / 14.5037738)
}

// Creates a pressure from bars.
func Bar(b float32) PoundPerSuqaredInch {
	return PoundPerSuqaredInch(b * 14.5037738)
}

// Main time unit of the protocol
type Second float32

// Gets the duration in seconds.
func (s Second) Second() float32 {
	return float32(s)
}

// Gets the elapsed time in [time.Duration].
func (s Second) Duration() time.Duration {
	return time.Duration(s) * 1000000000
}

// Main timestamp unit of the protocol
type MilliSecond uint32

// Gets the difference between two timestamps in milliseconds.
func (ms MilliSecond) MilliSecondsSince(ms2 MilliSecond) uint32 {
	return uint32(ms - ms2)
}

// Gets the difference between two timestamps in [time.Duration].
func (ms MilliSecond) DurationSince(ms2 MilliSecond) time.Duration {
	return time.Duration(ms.MilliSecondsSince(ms2)) * 1000000
}
