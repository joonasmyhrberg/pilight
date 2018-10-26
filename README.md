# pilight

This is a simple HomeKit application for white spectrum LED lamp control. It's designed to be used with two light strips, but anything dimmable with a MOSFET (such as IRLZ34N) should work.

It will work as-is on Raspberry Pi 3 when the cool light to is connected to pin 12 and the warm light to pin 35.

## Installing

Cross compiling with Raspbian: `GOARM=7 GOARCH=arm GOOS=linux go build`

For the PWM to work, `pilight` must be run as sudo.

