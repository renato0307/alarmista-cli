package main

import (
	"flag"
	"fmt"
	"os"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/raspi"
)

// readCharacteristic reads the value of a characteristic
func readCharacteristic(bleAdaptor *ble.ClientAdaptor, uuid string) {
	data, err := bleAdaptor.ReadCharacteristic(uuid)
	if err != nil {
		fmt.Println("error reading characteristic", err)
	} else {
		fmt.Println(uuid, "characteristic value is", string(data))
	}
}

// writeCharacteristic writes the value of a characteristic
func writeCharacteristic(bleAdaptor *ble.ClientAdaptor, uuid string, data []byte) {
	err := bleAdaptor.WriteCharacteristic(uuid, data)
	if err != nil {
		fmt.Println("error reading characteristic", err)
	} else {
		fmt.Println(uuid, "characteristic value was written")
	}
}

func main() {

	// read characteristic BLE command
	readCharacteristicCmd := flag.NewFlagSet("readc", flag.ExitOnError)
	readCAddress := readCharacteristicCmd.String("address", "", "the address of the BLE device, e.g. 30:AE:A4:02:BC:3A (required)")
	readCUuid := readCharacteristicCmd.String("uuid", "", "the UUID of the characteristic, e.g. cc0bd427-c9c3-43b0-a7c6-2df108b2b7c4 (required)")

	// write characteristic BLE command
	writeCharacteristicCmd := flag.NewFlagSet("writec", flag.ExitOnError)
	writeCAddress := writeCharacteristicCmd.String("address", "", "the address of the BLE device, e.g. 30:AE:A4:02:BC:3A (required)")
	writeCUuid := writeCharacteristicCmd.String("uuid", "", "the UUID of the characteristic, e.g. cc0bd427-c9c3-43b0-a7c6-2df108b2b7c4 (required)")
	writeCValue := writeCharacteristicCmd.String("value", "", "the value to write")

	// write gpio command
	gpioWriteCmd := flag.NewFlagSet("gpio", flag.ExitOnError)
	gpioWritePin := gpioWriteCmd.String("pin", "", "the pin number to write, e.g. 12 (required)")
	gpioWriteValue := gpioWriteCmd.Int("value", 1, "1 or 0")

	if len(os.Args) < 2 {
		fmt.Println("readc, writec or gpio subcommand is required")
		os.Exit(1)
	}

	// check which subcommand was executed
	switch os.Args[1] {
	case "readc":
		readCharacteristicCmd.Parse(os.Args[2:])
	case "writec":
		writeCharacteristicCmd.Parse(os.Args[2:])
	case "gpio":
		gpioWriteCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// handles read characteristic
	if readCharacteristicCmd.Parsed() {
		if *readCAddress == "" || *readCUuid == "" {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
			readCharacteristicCmd.PrintDefaults()
			os.Exit(1)
		}

		bleAdaptor := ble.NewClientAdaptor(*readCAddress)

		cerr := bleAdaptor.Connect()
		if cerr != nil {
			fmt.Println("connect NOK", cerr)
			return
		}
		fmt.Println("connect OK")

		readCharacteristic(bleAdaptor, *readCUuid)

		derr := bleAdaptor.Disconnect()
		if derr != nil {
			fmt.Println("disconnect NOK", derr)
		} else {
			fmt.Println("disconnect OK")
		}
	}

	// handles write characteristic
	if writeCharacteristicCmd.Parsed() {
		if *writeCAddress == "" || *writeCUuid == "" {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
			writeCharacteristicCmd.PrintDefaults()
			os.Exit(1)
		}

		bleAdaptor := ble.NewClientAdaptor(*writeCAddress)

		cerr := bleAdaptor.Connect()
		if cerr != nil {
			fmt.Println("connect NOK", cerr)
			return
		}
		fmt.Println("connect OK")

		writeCharacteristic(bleAdaptor, *writeCUuid, []byte(*writeCValue))

		derr := bleAdaptor.Disconnect()
		if derr != nil {
			fmt.Println("disconnect NOK", derr)
		} else {
			fmt.Println("disconnect OK")
		}
	}

	// handle gpio write
	if gpioWriteCmd.Parsed() {
		if *gpioWritePin == "" {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
			gpioWriteCmd.PrintDefaults()
			os.Exit(1)
		}

		fmt.Println("using gpio on pin", *gpioWritePin, "to send", *gpioWriteValue)

		r := raspi.NewAdaptor()
		writePin := gpio.NewDirectPinDriver(r, *gpioWritePin)
		if *gpioWriteValue == 1 {
			gerr := writePin.On()
			if gerr != nil {
				fmt.Println("on NOK", gerr)
				return
			}
			fmt.Println("on OK")
		} else if *gpioWriteValue == 0 {
			gerr := writePin.Off()
			if gerr != nil {
				fmt.Println("off NOK", gerr)
				return
			}
			fmt.Println("off OK")
		}
	} else {
		fmt.Println("gpio not parsed")
	}

}
