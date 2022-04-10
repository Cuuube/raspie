package main

import (
	"fmt"
	"raspie/gosrc/util"
	"time"
)

// const TEMPERTURE_PATH = "./test"

func main() {
	for {
		temperture, err := util.ReadTemperture()
		if err != nil {
			fmt.Println("硬件有误，检测终止！", err)
			return
		}
		fmt.Println("当前气温为：", temperture, "摄氏度。")

		// 一分钟后继续
		time.Sleep(time.Minute)
	}

}
