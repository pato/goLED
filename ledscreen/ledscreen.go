package main

import (
	led "github.com/pato/LEDserial/ledcomm"
	"github.com/vova616/screenshot"
	"image"
	"io"
	"time"
)

func toSimple(c uint32) uint8 {
	return uint8(c / 257)
}

func process(strip io.ReadWriteCloser) {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}

	bucketWidth := img.Bounds().Dx() / 60

	for i := 0; i < 60; i++ {
		r, g, b := extractColor(img, uint32(i*bucketWidth), uint32(bucketWidth))
		led.SetRGB(strip, uint8(i), r, g, b)
		time.Sleep(1 * time.Millisecond)
	}
	led.Flush(strip)
	time.Sleep(5 * time.Millisecond)
}

func extractColor(img *image.RGBA, start, width uint32) (uint8, uint8, uint8) {
	rSum, gSum, bSum := uint32(0), uint32(0), uint32(0)

	for i := start; i < start+width; i++ {
		r, g, b, _ := img.At(int(i), 540).RGBA()
		rSum += r
		gSum += g
		bSum += b
	}

	rAvg := rSum / width
	gAvg := gSum / width
	bAvg := bSum / width
	return toSimple(rAvg), toSimple(gAvg), toSimple(bAvg)
}

func main() {
	strip := led.Setup()
	led.Clear(strip)

	for {
		process(strip)
	}
}
