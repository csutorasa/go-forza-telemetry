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

// Byte order of the protocol.
var protocolByteOrder binary.ByteOrder = binary.LittleEndian

// Reads a uint8 at the index.
func parseUint8[T ~uint8](p Packet, i int) T {
	return (T)(p[i])
}

// Reads a uint16 at the index.
func parseUint16[T ~uint16](p Packet, i int) T {
	return (T)(protocolByteOrder.Uint16(p[i : i+2]))
}

// Reads a uint32 at the index.
func parseUint32[T ~uint32](p Packet, i int) T {
	return (T)(protocolByteOrder.Uint32(p[i : i+4]))
}

// Reads an int8 at the index.
func parseInt8[T ~int8](p Packet, i int) T {
	return (T)(parseUint8[uint8](p, i))
}

// Reads an int32 at the index.
func parseInt32[T ~int32](p Packet, i int) T {
	return (T)(parseUint32[uint32](p, i))
}

// Reads an int32 at the index, and converts it to bool.
func parseBool(p Packet, i int) bool {
	return parseInt32[int32](p, i) > 0
}

// Reads a float32 at the index.
func parseFloat32[T ~float32](p Packet, i int) T {
	return (T)(math.Float32frombits(parseUint32[uint32](p, i)))
}

// Reads 3 float32s at the index.
func parseThreeDimensionalFloat32[T ~float32](p Packet, i int) ThreeDimensional[T] {
	return ThreeDimensional[T]{
		X: parseFloat32[T](p, i),
		Y: parseFloat32[T](p, i+4),
		Z: parseFloat32[T](p, i+8),
	}
}

// Reads 4 int32s at the index, and converts them to bools.
func parseWheelBool(p Packet, i int) Wheel[bool] {
	return Wheel[bool]{
		FrontLeft:  parseBool(p, i),
		FrontRight: parseBool(p, i+4),
		RearLeft:   parseBool(p, i+8),
		RearRight:  parseBool(p, i+12),
	}
}

// Reads 4 float32 at the index.
func parseWheelFloat32[T ~float32](p Packet, i int) Wheel[T] {
	return Wheel[T]{
		FrontLeft:  parseFloat32[T](p, i),
		FrontRight: parseFloat32[T](p, i+4),
		RearLeft:   parseFloat32[T](p, i+8),
		RearRight:  parseFloat32[T](p, i+12),
	}
}

// Generic unmarshal function, that tries all versions to unmarshal.
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
