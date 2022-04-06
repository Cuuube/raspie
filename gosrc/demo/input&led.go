package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	PIN_INPUT = 22
	PIN_LED   = 4
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("初始化出错")
		return
	}
	defer rpio.Close()

	inputPin := rpio.Pin(PIN_INPUT)
	ledPin := rpio.Pin(PIN_LED)
	inputPin.Input()
	ledPin.Output()

	// init
	ledPin.Low()

	// 监听开关
	var on rpio.State
	go func() {
		for {
			on = inputPin.Read()
			time.Sleep(time.Millisecond * 10)
		}
	}()

	// 决定灯是否开或者关
	oldState := on
	go func() {
		for {
			if oldState != on {
				ledPin.Write(on)
				oldState = on
			}
			time.Sleep(time.Millisecond * 10)
		}
	}()

	// 10秒之后程序关闭
	select {
	case <-time.After(time.Second * 10):
		return
	}
}
