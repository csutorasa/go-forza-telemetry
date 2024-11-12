package forzatelemetry

const v2PacketSize = 324

// Forza Horizon 4 and 5 compatible Unmarshal
func UnmarshalV2(p Packet) (*ForzaData, error) {
	if len(p) != v2PacketSize {
		return nil, ErrInvalidPacketSize
	}
	return &ForzaData{
		IsRaceOn:                   p.Bool(0),
		TimestampMS:                p.Uint32(4),
		EngineMaxRpm:               p.Float32(8),
		EngineIdleRpm:              p.Float32(12),
		CurrentEngineRpm:           p.Float32(16),
		Acceleration:               p.ThreeDimensionalFloat32(20),
		Velocity:                   p.ThreeDimensionalFloat32(32),
		AngularVelocity:            p.ThreeDimensionalFloat32(44),
		Yaw:                        p.Float32(56),
		Pitch:                      p.Float32(60),
		Roll:                       p.Float32(64),
		NormalizedSuspensionTravel: p.WheelFloat32(68),
		TireSlipRatio:              p.WheelFloat32(84),
		WheelRotationSpeed:         p.WheelFloat32(100),
		WheelOnRumbleStrip:         p.WheelBool(116),
		WheelInPuddleDepth:         p.WheelFloat32(132),
		SurfaceRumble:              p.WheelFloat32(148),
		TireSlipAngle:              p.WheelFloat32(164),
		TireCombinedSlip:           p.WheelFloat32(180),
		SuspensionTravelMeters:     p.WheelFloat32(196),
		CarOrdinal:                 p.Int32(212),
		CarClass:                   (CarClass)(p.Int32(216)),
		CarPerformanceIndex:        p.Int32(220),
		DrivetrainType:             (DrivetrainType)(p.Int32(224)),
		NumCylinders:               p.Int32(228),
		// 12 gap
		Position:                    p.ThreeDimensionalFloat32(244),
		Speed:                       p.Float32(256),
		Power:                       p.Float32(260),
		Torque:                      p.Float32(264),
		TireTemp:                    p.WheelFloat32(268),
		Boost:                       p.Float32(284),
		Fuel:                        p.Float32(288),
		DistanceTraveled:            p.Float32(292),
		BestLap:                     p.Float32(296),
		LastLap:                     p.Float32(300),
		CurrentLap:                  p.Float32(304),
		CurrentRaceTime:             p.Float32(308),
		LapNumber:                   p.Uint16(312),
		RacePosition:                p.Uint8(314),
		Accel:                       p.Uint8(315),
		Brake:                       p.Uint8(316),
		Clutch:                      p.Uint8(317),
		HandBrake:                   p.Uint8(318),
		Gear:                        p.Uint8(319),
		Steer:                       p.Int8(320),
		NormalizedDrivingLine:       p.Int8(321),
		NormalizedAIBrakeDifference: p.Int8(322),
	}, nil
}
