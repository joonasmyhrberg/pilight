package main

import (
	"log"
	"os"

	"github.com/brutella/hc"
	hcAccessory "github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/myyra/pilight/accessory"
	rpio "github.com/stianeikeland/go-rpio"
)

const (
	minColorTemperature = 154
	maxColorTemperature = 370
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
	ledStrip.WhiteSpectrumBulb.ColorTemperature.SetMinValue(minColorTemperature)
	ledStrip.WhiteSpectrumBulb.ColorTemperature.SetMaxValue(maxColorTemperature)
	ledStrip.WhiteSpectrumBulb.ColorTemperature.SetValue(minColorTemperature)
	ledStrip.WhiteSpectrumBulb.On.SetValue(false)

	config := hc.Config{}
	t, err := hc.NewIPTransport(config, ledStrip.Accessory)
	if err != nil {
		log.Panic(err)
	}

	ledStrip.WhiteSpectrumBulb.On.OnValueRemoteUpdate(func(on bool) {
		setDuty(ledStrip, &coolLEDPin, &warmLEDPin)
	})

	ledStrip.WhiteSpectrumBulb.Brightness.OnValueRemoteUpdate(func(newBrightness int) {
		setDuty(ledStrip, &coolLEDPin, &warmLEDPin)
	})

	ledStrip.WhiteSpectrumBulb.ColorTemperature.OnValueRemoteUpdate(func(colorTemperature int) {
		setDuty(ledStrip, &coolLEDPin, &warmLEDPin)
	})

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}

func setDuty(bulb *accessory.WhiteSpectrumBulb, coolLEDPin *rpio.Pin, warmLEDPin *rpio.Pin) {

	if bulb.WhiteSpectrumBulb.On.GetValue() {

		brightnessMultiplier := bulb.WhiteSpectrumBulb.Brightness.GetValue()
		warmDutyMultiplier := getBalance(bulb.WhiteSpectrumBulb.ColorTemperature)
		coolDutyMultiplier := 1.0 - warmDutyMultiplier

		warmDuty := 100 * warmDutyMultiplier * float64(brightnessMultiplier) / 100.0
		coolDuty := 100 * coolDutyMultiplier * float64(brightnessMultiplier) / 100.0

		warmLEDPin.DutyCycle(uint32(warmDuty), 100)
		coolLEDPin.DutyCycle(uint32(coolDuty), 100)
	} else {
		coolLEDPin.DutyCycle(0, 100)
		warmLEDPin.DutyCycle(0, 100)
	}
}

// getBalance maps the possible colorBalance values to a number between 0 and 1.
// This is the multiplier for the warm LED brightness.
func getBalance(colorTemperature *characteristic.ColorTemperature) float64 {

	min := colorTemperature.GetMinValue()
	max := colorTemperature.GetMaxValue()
	current := colorTemperature.GetValue()
	return float64(current-min) / float64(max-min)
}
