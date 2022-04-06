package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("初始化出错")
		return
	}
	defer rpio.Close()

	pin := rpio.Pin(4)
	pin.Output()

	i := 0
	for i < 5 {
		pin.High()
		time.Sleep(time.Second)
		pin.Low()
		time.Sleep(time.Second)

		i++
	}
}
