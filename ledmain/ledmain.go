package main

import (
	"flag"
	"fmt"
	"github.com/pato/goLED/ledcomm"
	"time"
)

type Hertz float64
type BPM float64

func demo1(strip ledcomm.Strip, brightness float64) {
	strip.Clear()
	for {
		for i := uint8(0); i < 60; i++ {
			strip.SetHSV(i, float64(i*6), 1, brightness)
			strip.Flush()
			time.Sleep(4 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
		for i := uint8(0); i < 60; i++ {
			strip.SetRGB(i, 0, 0, 0)
			strip.Flush()
			time.Sleep(4 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
		for i := uint8(0); i < 60; i++ {
			strip.SetHSV(59-i, float64((59-i)*6), 1, brightness)
			strip.Flush()
			time.Sleep(4 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
		for i := uint8(0); i < 60; i++ {
			strip.SetRGB(59-i, 0, 0, 0)
			strip.Flush()
			time.Sleep(4 * time.Millisecond)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func demo2(strip ledcomm.Strip, brightness float64) {
	strip.Clear()
	var color float64 = 0
	for {
		for col := uint8(60); col > 0; col-- {
			for i := uint8(0); i < col; i++ {
				strip.SetHSV(i, color, 1, brightness)
				if i > 0 {
					strip.SetRGB(i-1, 0, 0, 0)
				}
				strip.Flush()
				time.Sleep(4 * time.Millisecond)
			}
		}
		color += 10
		if color > 359 {
			color = 10
		}
	}
}

func demo3(strip ledcomm.Strip, brightness float64) {
	strip.Clear()
	var colorStep float64 = 60
	var color float64 = 0
	var pastColor float64 = 59
	for {
		for col := uint8(60); col > 0; col-- {
			for i := uint8(0); i < col; i++ {
				strip.SetHSV(i, color, 1, brightness)
				if i > 0 {
					strip.SetHSV(i-1, pastColor, 1, brightness)
				}
				strip.Flush()
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

func demo4(strip ledcomm.Strip, brightness float64) {
	strip.Clear()
	for i := uint8(0); i < 60; i++ {
		strip.SetRGB(i, 255, 0, 0)
	}
	var color uint = 0
	var colorStep uint = 20
	for {
		for col := uint8(60); col > 0; col-- {
			for i := uint8(0); i < col; i++ {
				strip.SetHSV(i, float64(color), 1, brightness)
				if i == col-1 {
					strip.SetRGB(i-4, 0, 0, 0)
					strip.SetRGB(i-3, 0, 0, 0)
					strip.SetRGB(i-2, 0, 0, 0)
					strip.SetRGB(i-1, 0, 0, 0)
				} else if i > 3 {
					strip.SetHSV(i-4, float64(color), 0, 0)
					strip.SetHSV(i-3, float64(color), 0.9, brightness*0.1)
					strip.SetHSV(i-2, float64(color), 0.95, brightness*0.2)
					strip.SetHSV(i-1, float64(color), 1, brightness*0.5)
				}
				strip.Flush()
				time.Sleep(50 * time.Millisecond)
				color = (color + colorStep) % 360
			}
		}
	}
}

func demo5(strip ledcomm.Strip, brightness float64) {
	strip.Clear()
	for {
		for color := 0; color < 360; color++ {
			setStripHSV(strip, float64(color), 1, brightness)
			strip.Flush()
			time.Sleep(50 * time.Millisecond)
		}
		for color := 359; color >= 0; color-- {
			setStripHSV(strip, float64(color), 1, brightness)
			strip.Flush()
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func demo6(strip ledcomm.Strip, brightness float64) {
	strip.Clear()
	for {
		for color := 0; color < 360; color += 30 {
			setStripHSV(strip, float64(color), 1, brightness)
			strip.Flush()
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func demo7(strip ledcomm.Strip, brightness float64) {
	strip.Clear()
	for {
		for color := 0; color < 360; color += 60 {
			sendHSVFromMiddle(strip, float64(color), 1, brightness, bpmToFrequency(140))
		}
	}
}

func sendHSVFromMiddle(strip ledcomm.Strip, h, s, brightness float64, frequency Hertz) {
	leds := uint8(60)
	half := leds / 2
	for i := uint8(0); i < 30; i++ {
		strip.SetHSV(half-i, h, s, brightness)
		strip.SetHSV(half+i, h, s, brightness)
		strip.Flush()
		time.Sleep(frequencyToDelay(30, frequency))
	}
}

func bpmToFrequency(bpm BPM) Hertz {
	return Hertz(bpm / 60.0)
}

func frequencyToDelay(leds uint, frequency Hertz) time.Duration {
	return time.Millisecond * time.Duration((float64(1000.0) / float64(frequency) / float64(leds)))
}

func setStripRGB(strip ledcomm.Strip, r, g, b uint8) {
	for led := uint8(0); led < 60; led++ {
		strip.SetRGB(led, r, g, b)
	}
}

func setStripHSV(strip ledcomm.Strip, h, s, v float64) {
	for led := uint8(0); led < 60; led++ {
		strip.SetHSV(led, h, s, v)
	}
}

func main() {

	clear := flag.Bool("clear", false, "clears the led strip")
	demo := flag.Bool("demo", false, "run a basic demo that shows the color spectrum on the strip")
	send := flag.Bool("send", false, "set to send either an rgb or hsv color to the strip")
	brightness := flag.Float64("brightness", 255, "the maximum brightness in the demos [0-255]")
	n := flag.Int("n", 1, "which demo to run (requires -demo)")
	i := flag.Int("i", -1, "the led index (or all leds if no index given)")
	r := flag.Int("r", -1, "the red value [0-255]")
	g := flag.Int("g", -1, "the green value [0-255]")
	b := flag.Int("b", -1, "the blue value [0-255]")
	h := flag.Float64("h", -1, "the hue [0-359]")
	s := flag.Float64("s", -1, "the saturation [0-1]")
	v := flag.Float64("v", -1, "the value [0-255]")

	fmt.Printf("bpmToFrequencyToDelay(120) = %s\n", frequencyToDelay(30, bpmToFrequency(120)))

	flag.Parse()

	strip := ledcomm.Open()

	time.Sleep(1 * time.Second)

	if *clear {
		for i := 0; i < 10000; i++ {
			strip.Clear()
		}
	} else if *demo {
		switch *n {
		case 1:
			demo1(strip, *brightness)
		case 2:
			demo2(strip, *brightness)
		case 3:
			demo3(strip, *brightness)
		case 4:
			demo4(strip, *brightness)
		case 5:
			demo5(strip, *brightness)
		case 6:
			demo6(strip, *brightness)
		case 7:
			demo7(strip, *brightness)
		}

	} else if *send {
		if *r >= 0 && *g >= 0 && *b >= 0 {
			if *i >= 0 {
				strip.SetRGB(uint8(*i), uint8(*r), uint8(*g), uint8(*b))
			} else {
				for l := uint8(0); l < 60; l++ {
					strip.SetRGB(l, uint8(*r), uint8(*g), uint8(*b))
				}
			}
		} else if *h >= 0 && *s >= 0 && *v >= 0 {
			if *i >= 0 {
				strip.SetHSV(uint8(*i), *h, *s, *v)
			} else {
				for l := uint8(0); l < 60; l++ {
					strip.SetHSV(l, *h, *s, *v)
				}
			}
		} else {
			fmt.Printf("RGB or HSV need to be specified. See ledmain -help for usage\n")
		}
		strip.Flush()
	} else {
		flag.Usage()
	}
}
