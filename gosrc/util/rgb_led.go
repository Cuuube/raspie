package util

import "github.com/stianeikeland/go-rpio/v4"

type RGBLed struct {
	r_pin rpio.Pin
	g_pin rpio.Pin
	b_pin rpio.Pin
}

func (led RGBLed) Reset() {
	led.r_pin.Low()
	led.g_pin.Low()
	led.b_pin.Low()
}

func (led RGBLed) LightR() {
	led.g_pin.Low()
	led.b_pin.Low()
	led.r_pin.High()
}
func (led RGBLed) LightG() {
	led.r_pin.Low()
	led.g_pin.High()
	led.b_pin.Low()
}
func (led RGBLed) LightB() {
	led.r_pin.Low()
	led.g_pin.Low()
	led.b_pin.High()
}
