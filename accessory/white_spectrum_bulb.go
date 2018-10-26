package accessory

import (
	"github.com/brutella/hc/accessory"
	"github.com/myyra/pilight/service"
)

// WhiteSpectrumBulb is a Lightbulb accessory with one WhiteSpectrumBulb service
type WhiteSpectrumBulb struct {
	*accessory.Accessory
	WhiteSpectrumBulb *service.WhiteSpectrumBulb
}

// NewWhiteLightbulb returns an light bulb accessory with oneWhiteSpectrumBulb service
func NewWhiteLightbulb(info accessory.Info) *WhiteSpectrumBulb {
	acc := WhiteSpectrumBulb{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.WhiteSpectrumBulb = service.NewWhiteSpectrumBulb()

	acc.WhiteSpectrumBulb.Brightness.SetValue(100)

	acc.AddService(acc.WhiteSpectrumBulb.Service)

	return &acc
}
