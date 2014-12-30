package ledcomm

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/tarm/goserial"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const BaudRate float64 = 115200
const secondsPerBit = 1 / BaudRate
const microsecondsPerBit = secondsPerBit * (1000000)
const microsecondsPerByte = 8 * microsecondsPerBit
const MsPerColor = microsecondsPerByte * 5
const MsPerFlush = microsecondsPerByte
const MsPerClear = microsecondsPerByte

// SetHSV will convert the HSV color to RGB and then send over serial
func SetHSV(strip io.ReadWriteCloser, index uint8, h, s, v float64) {
	c := colorful.Hsv(h, s, v)
	SetRGB(strip, index, uint8(c.R), uint8(c.G), uint8(c.B))
}

// SetRGB will transfer the color to the correct index over serial
func SetRGB(s io.ReadWriteCloser, index, r, g, b uint8) {
	write(s, []byte{'s', r, g, b, index})
	time.Sleep(347 * time.Microsecond)
}

// Clear will send a clear signal over serial
func Clear(s io.ReadWriteCloser) {
	write(s, []byte{'c'})
	time.Sleep(69 * time.Microsecond)
}

// Flush will send a flush signal over serial
func Flush(s io.ReadWriteCloser) {
	write(s, []byte{'f'})
	time.Sleep(2150 * time.Microsecond)
}

func write(s io.ReadWriteCloser, data []byte) {
	_, err := s.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func ttyName() string {
	dev, err := ioutil.ReadDir("/dev")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dev {
		if strings.Contains(file.Name(), "ttyACM") {
			return "/dev/" + file.Name()
		}
	}

	log.Fatal("Could not find an appropriate tty connection")
	return ""
}

// SetupManual will initialize a serial connection to
// specified port at buad 115200 and return an
// io.ReadWriteCloser object to make further writes
func SetupManual(name string) io.ReadWriteCloser {
	c := &serial.Config{Name: name, Baud: 115200}
	strip, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return strip
}

// Setup will try to find the correct serial connection
// and then initialize the connection
func Setup() io.ReadWriteCloser {
	return SetupManual(ttyName())
}
