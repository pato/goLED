package main

import (
	"github.com/tarm/goserial"
	"io"
	"log"
	"time"
)

func setRGB(s io.ReadWriteCloser, index, r, g, b uint8) {
	write(s, []byte{'s', r, g, b, index})
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
}
