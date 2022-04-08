package main

import (
	"fmt"
	"time"

	"raspie/gosrc/util"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("初始化出错")
		return
	}
	defer rpio.Close()

	pin := rpio.Pin(util.GPIO0)
	pin.Input()

	i := 0
	for i < 5 {
		time.Sleep(time.Second)
		fmt.Println(pin.Read())

		i++
	}
}
