package forzatelemetry

type DrivetrainType int32

const (
	DrivetrainTypeFWD DrivetrainType = 0
	DrivetrainTypeRWD DrivetrainType = 1
	DrivetrainTypeAWD DrivetrainType = 2
)

type CarClass int32

const (
	CarClassD  CarClass = 0
	CarClassC  CarClass = 1
	CarClassB  CarClass = 2
	CarClassA  CarClass = 3
	CarClassS1 CarClass = 4
	CarClassS2 CarClass = 5
	CarClassX  CarClass = 6
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
	TimestampMS uint32
	// Maximum engine speed (1/min).
	EngineMaxRpm float32
	// Idle engine speed (1/min).
	EngineIdleRpm float32
	// Current engine speed (1/min).
	CurrentEngineRpm float32
	// Acceleration in car's local space (m/s2).
	Acceleration ThreeDimensional[float32]
	// Velocity in car's local space (m/s).
	Velocity ThreeDimensional[float32]
	// Angular velocity in car's local space (rad/s).
	AngularVelocity ThreeDimensional[float32]
	// Yaw in global space (rad).
	Yaw float32
	// Pitch in global space (rad).
	Pitch float32
	// Roll in global space (rad).
	Roll float32
	// Susepension travel (0-1 max strech to max compression).
	NormalizedSuspensionTravel Wheel[float32]
	// Tire slip ratio (above 1 means loss of grip).
	TireSlipRatio Wheel[float32]
	// Wheel roration speed (rad/s)
	WheelRotationSpeed Wheel[float32]
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
	SuspensionTravelMeters Wheel[float32]
	// Unique ID of the car make/model
	CarOrdinal int32
	// Car class
	CarClass CarClass
	// Performance index (100-999 worst to best car)
	CarPerformanceIndex int32
	// Drivetrain type
	DrivetrainType DrivetrainType
	// Number of cylinders
	NumCylinders int32
	// Car global position.
	Position ThreeDimensional[float32]
	// Car speed (m/s).
	Speed float32
	// Power (W).
	Power float32
	// Torque (Nm).
	Torque float32
	// Tire temperature (F).
	TireTemp Wheel[float32]
	// Turbo boost (Psi).
	Boost float32
	// Remaining fuel (0-1 empty to full).
	Fuel float32
	// Distance traveled during current race (m).
	DistanceTraveled float32
	// Best lap time (s).
	BestLap float32
	// Last lap time (s).
	LastLap float32
	// Current lap time (s).
	CurrentLap float32
	// Current race time (s).
	CurrentRaceTime float32
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
