# Changelog

## v1.2.0

Fix `Second.Duration()`.
Replace `MilliSecond.MilliSecond()` with `MilliSecond.MilliSecondsSince()`.
Replace `MilliSecond.Duration()` with `MilliSecond.DurationSince()`.
Add new creators for units from converted values.

## v1.1.0

`PerformanceIndex` is added.

Units are added with some conversion methods:

- Fahrenheit - Celsius
- Meter - centimeter, kilometer, inch, mile
- Meter/Second - kilometer/hour, mile/hour
- Radian - degree
- Watt - kilowatt, horsepower
- Newtonmeter - footpound
- Pound/SuqaredInch - bar
- Second, Millisecond - `time.Duration`

## v1.0.0

Initial release
