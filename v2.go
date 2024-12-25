package forzatelemetry

const v2PacketSize = 324

// Forza Horizon 4 and 5 compatible Unmarshal
func UnmarshalV2(p Packet) (*ForzaData, error) {
	if len(p) != v2PacketSize {
		return nil, ErrInvalidPacketSize
	}
	return &ForzaData{
		IsRaceOn:                   parseBool(p, 0),
		TimestampMS:                parseUint32[MilliSecond](p, 4),
		EngineMaxRpm:               parseFloat32[float32](p, 8),
		EngineIdleRpm:              parseFloat32[float32](p, 12),
		CurrentEngineRpm:           parseFloat32[float32](p, 16),
		Acceleration:               parseThreeDimensionalFloat32[MeterPerSecond](p, 20),
		Velocity:                   parseThreeDimensionalFloat32[MeterPerSecond](p, 32),
		AngularVelocity:            parseThreeDimensionalFloat32[Radian](p, 44),
		Yaw:                        parseFloat32[Radian](p, 56),
		Pitch:                      parseFloat32[Radian](p, 60),
		Roll:                       parseFloat32[Radian](p, 64),
		NormalizedSuspensionTravel: parseWheelFloat32[float32](p, 68),
		TireSlipRatio:              parseWheelFloat32[float32](p, 84),
		WheelRotationSpeed:         parseWheelFloat32[Radian](p, 100),
		WheelOnRumbleStrip:         parseWheelBool(p, 116),
		WheelInPuddleDepth:         parseWheelFloat32[float32](p, 132),
		SurfaceRumble:              parseWheelFloat32[float32](p, 148),
		TireSlipAngle:              parseWheelFloat32[float32](p, 164),
		TireCombinedSlip:           parseWheelFloat32[float32](p, 180),
		SuspensionTravelMeters:     parseWheelFloat32[Meter](p, 196),
		CarOrdinal:                 parseInt32[int32](p, 212),
		CarClass:                   parseInt32[CarClass](p, 216),
		CarPerformanceIndex:        parseInt32[PerformanceIndex](p, 220),
		DrivetrainType:             parseInt32[DrivetrainType](p, 224),
		NumCylinders:               parseInt32[int32](p, 228),
		// 12 gap
		Position:                    parseThreeDimensionalFloat32[Meter](p, 244),
		Speed:                       parseFloat32[MeterPerSecond](p, 256),
		Power:                       parseFloat32[Watt](p, 260),
		Torque:                      parseFloat32[NewtonMeter](p, 264),
		TireTemp:                    parseWheelFloat32[Fahrenheit](p, 268),
		Boost:                       parseFloat32[PoundPerSuqaredInch](p, 284),
		Fuel:                        parseFloat32[float32](p, 288),
		DistanceTraveled:            parseFloat32[Meter](p, 292),
		BestLap:                     parseFloat32[Second](p, 296),
		LastLap:                     parseFloat32[Second](p, 300),
		CurrentLap:                  parseFloat32[Second](p, 304),
		CurrentRaceTime:             parseFloat32[Second](p, 308),
		LapNumber:                   parseUint16[uint16](p, 312),
		RacePosition:                parseUint8[uint8](p, 314),
		Accel:                       parseUint8[uint8](p, 315),
		Brake:                       parseUint8[uint8](p, 316),
		Clutch:                      parseUint8[uint8](p, 317),
		HandBrake:                   parseUint8[uint8](p, 318),
		Gear:                        parseUint8[uint8](p, 319),
		Steer:                       parseInt8[int8](p, 320),
		NormalizedDrivingLine:       parseInt8[int8](p, 321),
		NormalizedAIBrakeDifference: parseInt8[int8](p, 322),
	}, nil
}
