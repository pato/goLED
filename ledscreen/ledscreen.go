package main

import (
	"github.com/BurntSushi/xgb"
	"github.com/pato/goLED/ledcomm"
	"github.com/pato/screenshot"
	"image"
)

func toSimple(c uint32) uint8 {
	return uint8(c / 257)
}

func process(xcon *xgb.Conn, strip ledcomm.Strip) {
	img, err := screenshot.CaptureScreen(xcon)
	if err != nil {
		panic(err)
	}

	bucketWidth := img.Bounds().Dx() / 60

	for i := 0; i < 60; i++ {
		r, g, b := extractColor(img, uint32(i*bucketWidth), uint32(bucketWidth))
		strip.SetRGB(uint8(i), r, g, b)
	}
	strip.Flush()
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
	strip := ledcomm.Open()
	xcon, err := screenshot.Setup()
	if err != nil {
		panic(err)
	}
	defer screenshot.Close(xcon)

	strip.Clear()

	for {
		process(xcon, strip)
	}
}
