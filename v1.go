package forzatelemetry

const v1PacketSize = 311

// Forza Motorsport 7 compatible Unmarshal
func UnmarshalV1(p Packet) (*ForzaData, error) {
	if len(p) != v1PacketSize {
		return nil, ErrInvalidPacketSize
	}
	return &ForzaData{
		IsRaceOn:                    p.Bool(0),
		TimestampMS:                 p.Uint32(4),
		EngineMaxRpm:                p.Float32(8),
		EngineIdleRpm:               p.Float32(12),
		CurrentEngineRpm:            p.Float32(16),
		Acceleration:                p.ThreeDimensionalFloat32(20),
		Velocity:                    p.ThreeDimensionalFloat32(32),
		AngularVelocity:             p.ThreeDimensionalFloat32(44),
		Yaw:                         p.Float32(56),
		Pitch:                       p.Float32(60),
		Roll:                        p.Float32(64),
		NormalizedSuspensionTravel:  p.WheelFloat32(68),
		TireSlipRatio:               p.WheelFloat32(84),
		WheelRotationSpeed:          p.WheelFloat32(100),
		WheelOnRumbleStrip:          p.WheelBool(116),
		WheelInPuddleDepth:          p.WheelFloat32(132),
		SurfaceRumble:               p.WheelFloat32(148),
		TireSlipAngle:               p.WheelFloat32(164),
		TireCombinedSlip:            p.WheelFloat32(180),
		SuspensionTravelMeters:      p.WheelFloat32(196),
		CarOrdinal:                  p.Int32(212),
		CarClass:                    (CarClass)(p.Int32(216)),
		CarPerformanceIndex:         p.Int32(220),
		DrivetrainType:              (DrivetrainType)(p.Int32(224)),
		NumCylinders:                p.Int32(228),
		Position:                    p.ThreeDimensionalFloat32(232),
		Speed:                       p.Float32(244),
		Power:                       p.Float32(248),
		Torque:                      p.Float32(252),
		TireTemp:                    p.WheelFloat32(256),
		Boost:                       p.Float32(272),
		Fuel:                        p.Float32(276),
		DistanceTraveled:            p.Float32(280),
		BestLap:                     p.Float32(284),
		LastLap:                     p.Float32(288),
		CurrentLap:                  p.Float32(292),
		CurrentRaceTime:             p.Float32(296),
		LapNumber:                   p.Uint16(300),
		RacePosition:                p.Uint8(302),
		Accel:                       p.Uint8(303),
		Brake:                       p.Uint8(304),
		Clutch:                      p.Uint8(305),
		HandBrake:                   p.Uint8(306),
		Gear:                        p.Uint8(307),
		Steer:                       p.Int8(308),
		NormalizedDrivingLine:       p.Int8(309),
		NormalizedAIBrakeDifference: p.Int8(310),
	}, nil
}
