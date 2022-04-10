package main

import (
	"fmt"
	"raspie/gosrc/util"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	MAX_RANGE = 255
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("初始化出错")
		return
	}
	defer rpio.Close()

	pin := rpio.Pin(util.GPIO24)
	// pin.Mode(rpio.Pwm)
	// pin.Freq(64000)
	// pin.DutyCycle(0, 32)

	// // the LED will be blinking at 2000Hz
	// // (source frequency divided by cycle length => 64000/32 = 2000)

	// // five times smoothly fade in and out
	// for i := 0; i < 5; i++ {
	// 	for i := uint32(0); i < 32; i++ { // increasing brightness
	// 		pin.DutyCycle(i, 32)
	// 		time.Sleep(time.Second / 32)
	// 	}
	// 	for i := uint32(32); i > 0; i-- { // decreasing brightness
	// 		pin.DutyCycle(i, 32)
	// 		time.Sleep(time.Second / 32)
	// 	}
	// }

	pin.Mode(rpio.Pwm)
	pin.Freq(64000)
	pin.DutyCycleWithPwmMode(0, 32, rpio.Balanced)
	// the LED will be blinking at 2000Hz
	// (source frequency divided by cycle length => 64000/32 = 2000)

	// five times smoothly fade in and out
	for i := 0; i < 5; i++ {
		for i := uint32(0); i < 32; i++ { // increasing brightness
			pin.DutyCycleWithPwmMode(i, 32, rpio.Balanced)
			time.Sleep(time.Second / 32)
		}
		for i := uint32(32); i > 0; i-- { // decreasing brightness
			pin.DutyCycleWithPwmMode(i, 32, rpio.Balanced)
			time.Sleep(time.Second / 32)
		}
	}

	pin.DutyCycleWithPwmMode(0, 32, rpio.Balanced)
}
