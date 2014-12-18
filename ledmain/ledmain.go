package main

import (
	"github.com/pato/LEDserial/ledcomm"
	"time"
)

func main() {
	strip := ledcomm.Setup()

	ledcomm.Clear(strip)

	for i := uint8(0); i < 60; i++ {
		ledcomm.SetHSV(strip, i, float64(i*6), 1, 55)
		time.Sleep(2 * time.Millisecond)
		ledcomm.Flush(strip)
		time.Sleep(4 * time.Millisecond)
	}
	ledcomm.Flush(strip)
}
