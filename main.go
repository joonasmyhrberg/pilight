package main

import (
	"log"
	"os"

	"github.com/brutella/hc"
	hcAccessory "github.com/brutella/hc/accessory"
	"github.com/joonasmyhrberg/pilight/accessory"
	rpio "github.com/stianeikeland/go-rpio"
)

func main() {

	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	coolLEDPin := rpio.Pin(18)
	warmLEDPin := rpio.Pin(19)
	coolLEDPin.Mode(rpio.Pwm)
	warmLEDPin.Mode(rpio.Pwm)
	// Only needed for one pin since they share the frequency
	coolLEDPin.Freq(1280000)
	coolLEDPin.DutyCycle(0, 100)
	warmLEDPin.DutyCycle(0, 100)

	info := hcAccessory.Info{
		Name:  "LEDStrip",
		Model: "PiLight",
	}

	ledStrip := accessory.NewWhiteLightbulb(info)
	ledStrip.WhiteSpectrumBulb.ColorTemperature.SetMinValue(154)
	ledStrip.WhiteSpectrumBulb.ColorTemperature.SetMaxValue(370)
	ledStrip.WhiteSpectrumBulb.ColorTemperature.SetValue(262)
	ledStrip.WhiteSpectrumBulb.On.SetValue(false)

	config := hc.Config{}
	t, err := hc.NewIPTransport(config, ledStrip.Accessory)
	if err != nil {
		log.Panic(err)
	}

	brightness := 0

	ledStrip.WhiteSpectrumBulb.On.OnValueRemoteUpdate(func(on bool) {
		if on {
			coolLEDPin.DutyCycle(uint32(brightness), 100)
			warmLEDPin.DutyCycle(uint32(brightness), 100)
		} else {
			coolLEDPin.DutyCycle(0, 100)
			warmLEDPin.DutyCycle(0, 100)
		}
	})

	ledStrip.WhiteSpectrumBulb.Brightness.OnValueRemoteUpdate(func(newBrightness int) {
		brightness = newBrightness
		coolLEDPin.DutyCycle(uint32(brightness), 100)
		warmLEDPin.DutyCycle(uint32(brightness), 100)
	})

	ledStrip.WhiteSpectrumBulb.ColorTemperature.OnValueRemoteUpdate(func(colorTemperature int) {
		log.Println("Color Temperature", colorTemperature)
	})

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
