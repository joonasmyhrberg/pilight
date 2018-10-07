package main

import (
	"log"

	"github.com/brutella/hc"
	hcAccessory "github.com/brutella/hc/accessory"
	"github.com/joonasmyhrberg/pilight/accessory"
)

func main() {

	info := hcAccessory.Info{
		Name:  "LEDStrip",
		Model: "PiLight",
	}

	ledStrip := accessory.NewWhiteLightbulb(info)
	ledStrip.WhiteSpectrumBulb.ColorTemperature.SetMinValue(154)
	ledStrip.WhiteSpectrumBulb.ColorTemperature.SetMaxValue(370)
	ledStrip.WhiteSpectrumBulb.ColorTemperature.SetValue(262)

	config := hc.Config{}
	t, err := hc.NewIPTransport(config, ledStrip.Accessory)
	if err != nil {
		log.Panic(err)
	}

	ledStrip.WhiteSpectrumBulb.On.OnValueRemoteUpdate(func(on bool) {
		log.Println("On", on)
	})

	ledStrip.WhiteSpectrumBulb.Brightness.OnValueRemoteUpdate(func(brightness int) {
		log.Println("Brightness", brightness)
	})

	ledStrip.WhiteSpectrumBulb.ColorTemperature.OnValueRemoteUpdate(func(colorTemperature int) {
		log.Println("Color Temperature", colorTemperature)
	})

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
