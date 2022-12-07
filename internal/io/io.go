package io

import (
	"machine"
	"time"
)

type PWMPeripheral interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Set(channel uint8, value uint32)
	Top() uint32
}

type DutyFn func(fraction float64) float64

type PulsingOutput struct {
	device    machine.Pin
	modulator PWMPeripheral
	period    time.Duration
	channel   uint8
	DutyFn    DutyFn
}

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

func (o *PulsingOutput) Set(fraction float64) {
	o.modulator.Set(o.channel, uint32(float64(o.modulator.Top())*o.DutyFn(fraction)))
}
