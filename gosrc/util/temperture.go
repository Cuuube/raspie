package util

import (
	"os"
	"regexp"
	"strconv"
)

const TEMPERTURE_PATH = "/sys/bus/w1/devices/28-01193e11f506/w1_slave"

func ReadTemperture() (float32, error) {
	bts, err := os.ReadFile(TEMPERTURE_PATH)
	if err != nil {
		return 0, err
	}
	reg, _ := regexp.Compile(`t=\d+`)
	slices := reg.FindAll(bts, 1)
	for _, slice := range slices {
		num, er := strconv.Atoi(string(slice[2:]))
		if er != nil {
			continue
		}
		return float32(num) / 1000, nil

	}
	return 0, nil
}
