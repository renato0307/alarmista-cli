package main

import (
        "fmt"
		"os"
		"flag"
        "gobot.io/x/gobot/platforms/ble"
)

func readCharacteristic(bleAdaptor *ble.ClientAdaptor, uuid string) {
	data, err := bleAdaptor.ReadCharacteristic(uuid)
	if err != nil {
		fmt.Println("Error reading characteristic", err)
	} else {
		fmt.Println(uuid, "characteristic value is", string(data))
	}
}

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])

	cerr := bleAdaptor.Connect()
	if cerr != nil {
		fmt.Println("Connect NOK", cerr);
		return
	}
	fmt.Println("Connect OK")

	readCharacteristic(bleAdaptor, "cc0bd427-c9c3-43b0-a7c6-2df108b2b7c4")
	readCharacteristic(bleAdaptor, "9200f537-e530-4f2a-b446-a5db5f3fc82f")

	derr := bleAdaptor.Disconnect()
	if derr != nil {
		fmt.Println("Disconnect NOK", derr);
	} else {
		fmt.Println("Disconnect OK");
	}

}