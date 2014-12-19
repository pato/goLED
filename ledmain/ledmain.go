package main

import (
	"flag"
	"fmt"
	"github.com/pato/LEDserial/ledcomm"
	"io"
	"time"
)

func demo1(strip io.ReadWriteCloser) {
	ledcomm.Clear(strip)
	for {
		for i := uint8(0); i < 60; i++ {
			ledcomm.SetHSV(strip, i, float64(i*6), 1, 55)
			ledcomm.Flush(strip)
			time.Sleep(4 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
		for i := uint8(0); i < 60; i++ {
			ledcomm.SetRGB(strip, i, 0, 0, 0)
			ledcomm.Flush(strip)
			time.Sleep(4 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
		for i := uint8(0); i < 60; i++ {
			ledcomm.SetHSV(strip, 59-i, float64((59-i)*6), 1, 55)
			ledcomm.Flush(strip)
			time.Sleep(4 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
		for i := uint8(0); i < 60; i++ {
			ledcomm.SetRGB(strip, 59-i, 0, 0, 0)
			ledcomm.Flush(strip)
			time.Sleep(4 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func demo2(strip io.ReadWriteCloser) {
	ledcomm.Clear(strip)
	var color float64 = 0
	var brightness float64 = 100
	for {
		for col := uint8(60); col > 0; col-- {
			for i := uint8(0); i < col; i++ {
				ledcomm.SetHSV(strip, i, color, 1, brightness)
				if i > 0 {
					ledcomm.SetRGB(strip, i-1, 0, 0, 0)
				}
				ledcomm.Flush(strip)
				time.Sleep(4 * time.Millisecond)
			}
		}
		color += 10
		if color > 359 {
			color = 10
		}
	}
}

func demo3(strip io.ReadWriteCloser) {
	ledcomm.Clear(strip)
	var brightness float64 = 100
	var colorStep float64 = 60
	var color float64 = 0
	var pastColor float64 = 59
	for {
		for col := uint8(60); col > 0; col-- {
			for i := uint8(0); i < col; i++ {
				ledcomm.SetHSV(strip, i, color, 1, brightness)
				if i > 0 {
					ledcomm.SetHSV(strip, i-1, pastColor, 1, brightness)
				}
				ledcomm.Flush(strip)
				time.Sleep(10 * time.Millisecond)
			}
		}
		pastColor = color
		color += colorStep
		if color > 359 {
			color = colorStep
		}
	}
}

func main() {

	clear := flag.Bool("clear", false, "clears the led strip")
	demo := flag.Bool("demo", false, "run a basic demo that shows the color spectrum on the strip")
	send := flag.Bool("send", false, "set to send either an rgb or hsv color to the strip")
	n := flag.Int("n", 1, "which demo to run (requires -demo)")
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
		switch *n {
		case 1:
			demo1(strip)
		case 2:
			demo2(strip)
		case 3:
			demo3(strip)
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
	} else {
		flag.Usage()
	}

}
