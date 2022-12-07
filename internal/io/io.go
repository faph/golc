package io

import (
	"machine"
	"time"
)

// PWMPeripheral is an interface for a Pulse Width Modulation device.
// The machine library does not export this type.
type PWMPeripheral interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Set(channel uint8, value uint32)
	Top() uint32
}

// DutyFn is a function type for a duty function to translate an input value into an output duty value
type DutyFn func(fraction float64) float64

// PulsingOutput is a high-level data type to setup a PWM output
type PulsingOutput struct {
	device    machine.Pin
	modulator PWMPeripheral
	period    time.Duration
	channel   uint8
	DutyFn    DutyFn
}

// Create a new PulsingOutput device
func NewPulsingOutput(device machine.Pin, modulator PWMPeripheral, period time.Duration) (*PulsingOutput, error) {
	modulator.Configure(machine.PWMConfig{
		Period: uint64(period.Nanoseconds()),
	})
	ch, err := modulator.Channel(device)
	if err != nil {
		return nil, err
	}
	result := &PulsingOutput{
		device:    device,
		modulator: modulator,
		period:    period,
		channel:   ch,
		DutyFn:    func(fraction float64) float64 { return fraction },
	}
	return result, nil
}

// Set the duty value
func (o *PulsingOutput) Set(fraction float64) {
	o.modulator.Set(o.channel, uint32(float64(o.modulator.Top())*o.DutyFn(fraction)))
}
