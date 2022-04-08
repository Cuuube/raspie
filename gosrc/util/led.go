package util

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func Slash(pin rpio.Pin) {
	pin.Low()
	pin.High()
	time.Sleep(time.Millisecond * 100)
	pin.Low()
}
