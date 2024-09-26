package main

import (
	"fmt"
)

func main() {
	// Discover peers
	peers := PeerDiscovery()

	// Ask the user to select a peer
	selectedName, selectedSocket := AskForPeerSelection(peers)
	fmt.Printf("You selected %s (%s)\n", selectedName, selectedSocket)

	// Ask the user to enter the path to the file to be sent
	filePath := AskForFilePath()

	// Send the file
	TransferFile(selectedSocket, filePath)
}
