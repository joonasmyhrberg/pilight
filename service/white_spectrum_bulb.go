package service

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

// WhiteSpectrumBulb defines a Lightbulb service for lamps that can only show different white temperatures.
type WhiteSpectrumBulb struct {
	*service.Service

	On               *characteristic.On
	Brightness       *characteristic.Brightness
	ColorTemperature *characteristic.ColorTemperature
}

// NewWhiteSpectrumBulb returns a WhiteSpectrumBulb
func NewWhiteSpectrumBulb() *WhiteSpectrumBulb {

	svc := WhiteSpectrumBulb{}
	svc.Service = service.New(service.TypeLightbulb)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	svc.ColorTemperature = characteristic.NewColorTemperature()
	svc.AddCharacteristic(svc.ColorTemperature.Characteristic)

	return &svc
}
