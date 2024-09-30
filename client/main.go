package main

import (
	"fmt"
)

func main() {
	// Read the path from the CLI
	filePath := ReadFilePathFromCLI()

	// Discover peers
	peers := PeerDiscovery()

	// Ask the user to select a peer
	selectedName, selectedSocket := AskForPeerSelection(peers)
	fmt.Printf("You selected %s (%s)\n", selectedName, selectedSocket)

	// Send the file
	TransferFile(selectedSocket, filePath)
}
