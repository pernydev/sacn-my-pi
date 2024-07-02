package main

import (
	"fmt"

	"github.com/Hundemeier/go-sacn/sacn"
	"github.com/charmbracelet/log"
)

// make a 512 channel DMX universe with all 0s

var universe = make([]byte, 512)

func main() {
	receiver := connectSACN()
	defer receiver.Close()

	/*dmx, err := dmx.NewDMX(512, "AL05ZDF6")

	if err != nil {
		log.Fatal(err)
	}
	defer dmx.Close()

	log.Info("Setting all channels to 0")
	*/

	receiver.SetOnChangeCallback(onPacketChange)
	receiver.Start()
	log.Info("Listening for sACN packets")

	fmt.Print("\n\n sACN My Pi active. Press enter to exit.\n\n")

	fmt.Scanln()
}

func connectSACN() *sacn.ReceiverSocket {
	// Create a new sACN receiver
	receiver, err := sacn.NewReceiverSocket("", nil)

	if err != nil {
		log.Fatal(err)
	}

	receiver.JoinUniverse(1)
	log.Info("Joined universe 1")
	return receiver
}

func onPacketChange(_ sacn.DataPacket, new sacn.DataPacket) {
	for i, channel := range new.Data() {
		universe[i] = channel
	}
	log.Info("Universe updated")
}
