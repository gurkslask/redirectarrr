package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
)

func main() {
	fmt.Println("vim-go")
	options := serial.OpenOptions{
		PortName:        "COM10",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatlf("serial.Open: %v", err)
	}
	defer port.Close()

}
