package main

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"log"
	"time"
)

func setRGB(s io.ReadWriteCloser, index, r, g, b uint8) {
	data := []byte{r, g, b, index}
	_, err := s.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func clear(s io.ReadWriteCloser) {
	for i := uint8(0); i < 60; i++ {
		setRGB(s, i, 0, 0, 0)
		time.Sleep(10000000)
		fmt.Printf("i: %d\n", i)
	}
}

func main() {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
	strip, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	clear(strip)

	//	setRGB(s, 0, 0, 255, 0)
	//	time.Sleep(1)
	//	setRGB(s, 0, 1, 255, 0)
	//	time.Sleep(1)
	//	setRGB(s, 0, 2, 255, 0)
}
