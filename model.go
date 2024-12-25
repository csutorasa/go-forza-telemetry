package forzatelemetry

// Type of car drvietrain, that shows which wheels are driven.
type DrivetrainType int32

const (
	// Front wheel drive
	DrivetrainTypeFWD DrivetrainType = 0
	// Rear wheel drive
	DrivetrainTypeRWD DrivetrainType = 1
	// All wheel drive
	DrivetrainTypeAWD DrivetrainType = 2
)

// Car performance index (100-999 worst to best car).
type PerformanceIndex int32

func (pi PerformanceIndex) Class() CarClass {
	if pi < 100 || pi > 999 {
		panic("unexpected performance index")
	}
	if pi <= 500 {
		return CarClassD
	}
	if pi <= 600 {
		return CarClassC
	}
	if pi <= 700 {
		return CarClassB
	}
	if pi <= 800 {
		return CarClassA
	}
	if pi <= 900 {
		return CarClassS1
	}
	if pi <= 998 {
		return CarClassS2
	}
	return CarClassX
}

// Car performance class
type CarClass int32

const (
	// 100-500 performance index
	CarClassD CarClass = 0
	// 501-600 performance index
	CarClassC CarClass = 1
	// 601-700 performance index
	CarClassB CarClass = 2
	// 701-800 performance index
	CarClassA CarClass = 3
	// 801-900 performance index
	CarClassS1 CarClass = 4
	// 901-998 performance index
	CarClassS2 CarClass = 5
	// 999 performance index
	CarClassX CarClass = 6
)

// Values that are three dimensional.
type ThreeDimensional[T any] struct {
	X T
	Y T
	Z T
}

// Values that are for each wheel.
type Wheel[T any] struct {
	FrontLeft  T
	FrontRight T
	RearLeft   T
	RearRight  T
}

// Parsed packet data.
type ForzaData struct {
	// True if the game is not paused.
	IsRaceOn bool
	// Milliseconds since the game has started.
	TimestampMS MilliSecond
	// Maximum engine speed (1/min).
	EngineMaxRpm float32
	// Idle engine speed (1/min).
	EngineIdleRpm float32
	// Current engine speed (1/min).
	CurrentEngineRpm float32
	// Acceleration in car's local space (m/s2).
	Acceleration ThreeDimensional[MeterPerSecond]
	// Velocity in car's local space (m/s).
	Velocity ThreeDimensional[MeterPerSecond]
	// Angular velocity in car's local space (rad/s).
	AngularVelocity ThreeDimensional[Radian]
	// Yaw in global space (rad).
	Yaw Radian
	// Pitch in global space (rad).
	Pitch Radian
	// Roll in global space (rad).
	Roll Radian
	// Susepension travel (0-1 max strech to max compression).
	NormalizedSuspensionTravel Wheel[float32]
	// Tire slip ratio (above 1 means loss of grip).
	TireSlipRatio Wheel[float32]
	// Wheel roration speed (rad/s)
	WheelRotationSpeed Wheel[Radian]
	// If the wheel is on rumble strip.
	WheelOnRumbleStrip Wheel[bool]
	// Wheel in puddle depth (0-1 not in puddle to in the deepest puddle).
	WheelInPuddleDepth Wheel[float32]
	// Non-dimensional surface rumble values passed to controller force feedback.
	SurfaceRumble Wheel[float32]
	// Tire normalized slip angle (above 1 means loss of grip).
	TireSlipAngle Wheel[float32]
	// Tire normalized combined slip (above 1 means loss of grip).
	TireCombinedSlip Wheel[float32]
	// Suspension travel (m).
	SuspensionTravelMeters Wheel[Meter]
	// Unique ID of the car make/model
	CarOrdinal int32
	// Car class
	CarClass CarClass
	// Performance index (100-999 worst to best car)
	CarPerformanceIndex PerformanceIndex
	// Drivetrain type
	DrivetrainType DrivetrainType
	// Number of cylinders
	NumCylinders int32
	// Car global position.
	Position ThreeDimensional[Meter]
	// Car speed (m/s).
	Speed MeterPerSecond
	// Power (W).
	Power Watt
	// Torque (Nm).
	Torque NewtonMeter
	// Tire temperature (F).
	TireTemp Wheel[Fahrenheit]
	// Turbo boost (Psi).
	Boost PoundPerSuqaredInch
	// Remaining fuel (0-1 empty to full).
	Fuel float32
	// Distance traveled during current race (m).
	DistanceTraveled Meter
	// Best lap time (s).
	BestLap Second
	// Last lap time (s).
	LastLap Second
	// Current lap time (s).
	CurrentLap Second
	// Current race time (s).
	CurrentRaceTime Second
	// Lap number.
	LapNumber uint16
	// Race position.
	RacePosition uint8
	// Accelerator position (0-255 released to full).
	Accel uint8
	// Brake position (0-255 released to full).
	Brake uint8
	// Clutch position (0-255 released to full).
	Clutch uint8
	// Handbrake position (0-255 released to full).
	HandBrake uint8
	// Gear
	Gear uint8
	// Steering wheel position (-128-127 left to right)
	Steer                       int8
	NormalizedDrivingLine       int8
	NormalizedAIBrakeDifference int8
}
