package main

import (
	"flag"
	"fmt"
	"github.com/pato/LEDserial/ledcomm"
	"time"
)

func main() {

	clear := flag.Bool("clear", false, "clears the led strip")
	demo := flag.Bool("demo", false, "run a basic demo that shows the color spectrum on the strip")
	send := flag.Bool("send", false, "set to send either an rgb or hsv color to the strip")
	i := flag.Int("i", 0, "the led index")
	r := flag.Int("r", -1, "the red value [0-255]")
	g := flag.Int("g", -1, "the green value [0-255]")
	b := flag.Int("b", -1, "the blue value [0-255]")
	h := flag.Float64("h", -1, "the hue [0-359]")
	s := flag.Float64("s", -1, "the saturation [0-1]")
	v := flag.Float64("v", -1, "the value [0-255]")

	flag.Parse()

	strip := ledcomm.Setup("/dev/ttyACM0")

	if *clear {
		ledcomm.Clear(strip)
	} else if *demo {
		ledcomm.Clear(strip)
		for {
			for i := uint8(0); i < 60; i++ {
				ledcomm.SetHSV(strip, i, float64(i*6), 1, 55)
				time.Sleep(347 * time.Microsecond)
				ledcomm.Flush(strip)
				time.Sleep(4 * time.Millisecond)
			}
			time.Sleep(100 * time.Millisecond)
			for i := uint8(0); i < 60; i++ {
				ledcomm.SetRGB(strip, i, 0, 0, 0)
				time.Sleep(347 * time.Microsecond)
				ledcomm.Flush(strip)
				time.Sleep(4 * time.Millisecond)
			}
			time.Sleep(100 * time.Millisecond)
			for i := uint8(0); i < 60; i++ {
				ledcomm.SetHSV(strip, 59-i, float64((59-i)*6), 1, 55)
				time.Sleep(347 * time.Microsecond)
				ledcomm.Flush(strip)
				time.Sleep(4 * time.Millisecond)
			}
			time.Sleep(100 * time.Millisecond)
			for i := uint8(0); i < 60; i++ {
				ledcomm.SetRGB(strip, 59-i, 0, 0, 0)
				time.Sleep(347 * time.Microsecond)
				ledcomm.Flush(strip)
				time.Sleep(4 * time.Millisecond)
			}
			time.Sleep(100 * time.Millisecond)
		}
	} else if *send {
		if *r >= 0 && *g >= 0 && *b >= 0 {
			ledcomm.SetRGB(strip, uint8(*i), uint8(*r), uint8(*g), uint8(*b))
		} else if *h >= 0 && *s >= 0 && *v >= 0 {
			ledcomm.SetHSV(strip, uint8(*i), *h, *s, *v)
		} else {
			fmt.Printf("RGB or HSV need to be specified. See ledmain -help for usage\n")
		}
		ledcomm.Flush(strip)
	}

}
