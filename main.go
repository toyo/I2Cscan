package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_100KHZ,
	})

	for {
		fmt.Println()
		fmt.Print(`    0  1  2  3  4  5  6  7  8  9  A  B  C  D  E  F`)
		for y := uint16(0); y < 0x8; y++ {
			fmt.Printf("\n%02x", y*0x10)
			for x := uint16(0); x < 0x10; x++ {
				i2caddr := y*0x10 + x
				err := machine.I2C0.Tx(i2caddr, []byte{0}, nil)
				if err == nil {
					fmt.Printf(" %02x", i2caddr)
				} else {
					fmt.Print(" --")
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
