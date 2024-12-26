package main

import (
	"fmt"

	forzatelemetry "github.com/csutorasa/go-forza-telemetry"
)

func main() {
	// Creates a new server.
	s, err := forzatelemetry.NewServerPort(12000)
	if err != nil {
		panic(err)
	}
	// Listen to packets.
	err = s.Listen(func(p forzatelemetry.Packet) {
		// Unmarshal incoming packet.
		d, err := forzatelemetry.Unmarshal(p)
		if err != nil {
			panic(err)
		}
		// Display data.
		if d.IsRaceOn {
			fmt.Printf("%f at %f RPM in %d gear\n", d.Speed, d.CurrentEngineRpm, d.Gear)
		}
	})
	if err != nil {
		panic(err)
	}
}
