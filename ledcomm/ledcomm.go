package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/tarm/goserial"
	"io"
	"log"
	"time"
)

// SetHSV will convert the HSV color to RGB and then send over serial
func SetHSV(strip io.ReadWriteCloser, index uint8, h, s, v float64) {
	c := colorful.Hsv(h, s, v)
	SetRGB(strip, index, uint8(c.R), uint8(c.G), uint8(c.B))
}

// SetRGB will transfer the color to the correct index over serial
func SetRGB(s io.ReadWriteCloser, index, r, g, b uint8) {
	write(s, []byte{'s', r, g, b, index})
}

// Clear will send a clear signal over serial
func Clear(s io.ReadWriteCloser) {
	write(s, []byte{'c'})
}

// Flush will send a flush signal over serial
func Flush(s io.ReadWriteCloser) {
	write(s, []byte{'f'})
}

func write(s io.ReadWriteCloser, data []byte) {
	_, err := s.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func Setup() io.ReadWriteCloser {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
	strip, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return strip
}

func demo() {
	strip := Setup()

	Clear(strip)

	for i := uint8(0); i < 60; i++ {
		SetHSV(strip, i, float64(i*6), 1, 55)
		time.Sleep(5 * time.Millisecond)
		Flush(strip)
		time.Sleep(10 * time.Millisecond)
	}
	Flush(strip)
}
