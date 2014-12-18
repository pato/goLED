package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/tarm/goserial"
	"io"
	"log"
	"time"
)

func setRGB(s io.ReadWriteCloser, index, r, g, b uint8) {
	write(s, []byte{'s', r, g, b, index})
}

func setHSV(strip io.ReadWriteCloser, index uint8, h, s, v float64) {
	c := colorful.Hsv(h, s, v)
	setRGB(strip, index, uint8(c.R), uint8(c.G), uint8(c.B))
}

func clear(s io.ReadWriteCloser) {
	write(s, []byte{'c'})
}

func flush(s io.ReadWriteCloser) {
	write(s, []byte{'f'})
}

func write(s io.ReadWriteCloser, data []byte) {
	_, err := s.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
	strip, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	clear(strip)

	for i := uint8(0); i < 60; i++ {
		setHSV(strip, i, float64(i*6), 1, 255)
		time.Sleep(5 * time.Millisecond)
		flush(strip)
		time.Sleep(10 * time.Millisecond)
	}
	flush(strip)
}
