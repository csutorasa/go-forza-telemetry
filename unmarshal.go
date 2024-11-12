package forzatelemetry

import (
	"encoding/binary"
	"fmt"
	"math"
)

// Error for showing incompatible data length.
var ErrInvalidPacketSize = fmt.Errorf("invalid packet size")

// Forza UDP packet data.
type Packet []byte

var protocolByteOrder binary.ByteOrder = binary.LittleEndian

// Reads a uint8 at the index.
func (p Packet) Uint8(i int) uint8 {
	return p[i]
}

// Reads a uint16 at the index.
func (p Packet) Uint16(i int) uint16 {
	return protocolByteOrder.Uint16(p[i : i+2])
}

// Reads a uint32 at the index.
func (p Packet) Uint32(i int) uint32 {
	return protocolByteOrder.Uint32(p[i : i+4])
}

// Reads an int8 at the index.
func (p Packet) Int8(i int) int8 {
	return (int8)(p.Uint8(i))
}

// Reads an int32 at the index.
func (p Packet) Int32(i int) int32 {
	return (int32)(p.Uint32(i))
}

// Reads an int32 at the index, and converts it to bool.
func (p Packet) Bool(i int) bool {
	return p.Int32(i) > 0
}

// Reads a float32 at the index.
func (p Packet) Float32(i int) float32 {
	return math.Float32frombits(p.Uint32(i))
}

// Reads 3 float32s at the index.
func (p Packet) ThreeDimensionalFloat32(i int) ThreeDimensional[float32] {
	return ThreeDimensional[float32]{
		X: p.Float32(i),
		Y: p.Float32(i + 4),
		Z: p.Float32(i + 8),
	}
}

// Reads 4 int32s at the index, and converts them to bools.
func (p Packet) WheelBool(i int) Wheel[bool] {
	return Wheel[bool]{
		FrontLeft:  p.Bool(i),
		FrontRight: p.Bool(i + 4),
		RearLeft:   p.Bool(i + 8),
		RearRight:  p.Bool(i + 12),
	}
}

// Reads 4 float32 at the index.
func (p Packet) WheelFloat32(i int) Wheel[float32] {
	return Wheel[float32]{
		FrontLeft:  math.Float32frombits(p.Uint32(i)),
		FrontRight: math.Float32frombits(p.Uint32(i + 4)),
		RearLeft:   math.Float32frombits(p.Uint32(i + 8)),
		RearRight:  math.Float32frombits(p.Uint32(i + 12)),
	}
}

// Generic unmarshal
func Unmarshal(p Packet) (*ForzaData, error) {
	d, err := UnmarshalV1(p)
	if err == nil {
		return d, nil
	}
	d, err = UnmarshalV2(p)
	if err == nil {
		return d, nil
	}
	d, err = UnmarshalV3(p)
	if err == nil {
		return d, nil
	}
	return nil, ErrInvalidPacketSize
}
