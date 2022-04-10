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

	pin := rpio.Pin(util.GPIO0)
	pin.Pwm()
	pin.Freq(60)
	pin.DutyCycle(0, MAX_RANGE)

	var i uint32 = 0
	for i < 255 {
		pin.DutyCycle(i, MAX_RANGE)

		time.Sleep(time.Second / 255)
		i++
	}
	for i >= 255 {
		pin.DutyCycle(i, MAX_RANGE)

		time.Sleep(time.Second / 255)
		i--
	}
}
