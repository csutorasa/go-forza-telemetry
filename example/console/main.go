package main

import (
	"fmt"

	forzatelemetry "github.com/csutorasa/go-forza-telemetry"
)

func main() {
	s, err := forzatelemetry.NewServerPort(12000)
	if err != nil {
		panic(err)
	}
	err = s.Listen(func(p forzatelemetry.Packet) {
		d, err := forzatelemetry.Unmarshal(p)
		if err != nil {
			panic(err)
		}
		if d.IsRaceOn {
			fmt.Printf("%f at %f RPM in %d gear\n", d.Speed, d.CurrentEngineRpm, d.Gear)
		}
	})
	if err != nil {
		panic(err)
	}
}
