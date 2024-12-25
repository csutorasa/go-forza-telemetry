package forzatelemetry

const v1PacketSize = 311

// Forza Motorsport 7 compatible Unmarshal
func UnmarshalV1(p Packet) (*ForzaData, error) {
	if len(p) != v1PacketSize {
		return nil, ErrInvalidPacketSize
	}
	return &ForzaData{
		IsRaceOn:                    parseBool(p, 0),
		TimestampMS:                 parseUint32[MilliSecond](p, 4),
		EngineMaxRpm:                parseFloat32[float32](p, 8),
		EngineIdleRpm:               parseFloat32[float32](p, 12),
		CurrentEngineRpm:            parseFloat32[float32](p, 16),
		Acceleration:                parseThreeDimensionalFloat32[MeterPerSecond](p, 20),
		Velocity:                    parseThreeDimensionalFloat32[MeterPerSecond](p, 32),
		AngularVelocity:             parseThreeDimensionalFloat32[Radian](p, 44),
		Yaw:                         parseFloat32[Radian](p, 56),
		Pitch:                       parseFloat32[Radian](p, 60),
		Roll:                        parseFloat32[Radian](p, 64),
		NormalizedSuspensionTravel:  parseWheelFloat32[float32](p, 68),
		TireSlipRatio:               parseWheelFloat32[float32](p, 84),
		WheelRotationSpeed:          parseWheelFloat32[Radian](p, 100),
		WheelOnRumbleStrip:          parseWheelBool(p, 116),
		WheelInPuddleDepth:          parseWheelFloat32[float32](p, 132),
		SurfaceRumble:               parseWheelFloat32[float32](p, 148),
		TireSlipAngle:               parseWheelFloat32[float32](p, 164),
		TireCombinedSlip:            parseWheelFloat32[float32](p, 180),
		SuspensionTravelMeters:      parseWheelFloat32[Meter](p, 196),
		CarOrdinal:                  parseInt32[int32](p, 212),
		CarClass:                    parseInt32[CarClass](p, 216),
		CarPerformanceIndex:         parseInt32[PerformanceIndex](p, 220),
		DrivetrainType:              parseInt32[DrivetrainType](p, 224),
		NumCylinders:                parseInt32[int32](p, 228),
		Position:                    parseThreeDimensionalFloat32[Meter](p, 232),
		Speed:                       parseFloat32[MeterPerSecond](p, 244),
		Power:                       parseFloat32[Watt](p, 248),
		Torque:                      parseFloat32[NewtonMeter](p, 252),
		TireTemp:                    parseWheelFloat32[Fahrenheit](p, 256),
		Boost:                       parseFloat32[PoundPerSuqaredInch](p, 272),
		Fuel:                        parseFloat32[float32](p, 276),
		DistanceTraveled:            parseFloat32[Meter](p, 280),
		BestLap:                     parseFloat32[Second](p, 284),
		LastLap:                     parseFloat32[Second](p, 288),
		CurrentLap:                  parseFloat32[Second](p, 292),
		CurrentRaceTime:             parseFloat32[Second](p, 296),
		LapNumber:                   parseUint16[uint16](p, 300),
		RacePosition:                parseUint8[uint8](p, 302),
		Accel:                       parseUint8[uint8](p, 303),
		Brake:                       parseUint8[uint8](p, 304),
		Clutch:                      parseUint8[uint8](p, 305),
		HandBrake:                   parseUint8[uint8](p, 306),
		Gear:                        parseUint8[uint8](p, 307),
		Steer:                       parseInt8[int8](p, 308),
		NormalizedDrivingLine:       parseInt8[int8](p, 309),
		NormalizedAIBrakeDifference: parseInt8[int8](p, 310),
	}, nil
}
