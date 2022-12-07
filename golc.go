package main

import (
	"machine"
	"math"
	"time"

	"github.com/faph/golc/internal/io"
)

func main() {
	pwm, err := io.NewPulsingOutput(machine.LED, machine.PWM4, 10*time.Millisecond)
	pwm.DutyFn = ledBrightness
	if err != nil {
		println(err)
		return
	}

	for {
		for i := float64(0); i <= 1; i += .001 {
			pwm.Set(i)
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func ledBrightness(fraction float64) float64 {
	return math.Min(math.Max(0, 1.008*fraction*fraction-0.1111*fraction), 1)
}
