package main

import (
	"file_transfer_naive/formatting"
	"file_transfer_naive/myutils"
	"fmt"
	"log"
)

// AskForPeerSelection asks the user to select a peer from the list of peers
func AskForPeerSelection(receiverIDs []formatting.ServerID) (receiverName string, receiverSocket string) {
	// Show the list of peers
	peers := append(receiverIDs, formatting.ServerID{
		PeerRespMsg: formatting.PeerRespMsg{
			Name:              "Custom",
			FileReceptionPort: "Custom",
		},
		ServerIp: "Custom",
	})
	toPrint := "Peers:\n"
	numbering := 1
	for _, peer := range peers {
		toPrint += fmt.Sprintf("%v Name: %s; Socket for receiving files: %s:%s\n",
			numbering, peer.Name, peer.ServerIp, peer.FileReceptionPort)
		numbering++
	}
	fmt.Println(toPrint)

	// Ask the user to select a peer
	fmt.Printf("Please enter the No. of the peer you want to send the file to:\n> ")
	var peerNo int
	nArgs, err := fmt.Scan(&peerNo)
	if err != nil {
		log.Fatalf("Error scanning input: %s\n", err)
	}
	if nArgs != 1 {
		log.Fatalf("Expected 1 argument, got %d\n", nArgs)
	}
	if peerNo < 1 || peerNo > len(peers) {
		log.Fatalf("Invalid peer number: %d\n", peerNo)
	}
	if peerNo == len(peers) {
		receiverName = "Custom"
		receiverSocket = AskForSocket()
		return
	}

	receiverId := receiverIDs[peerNo-1]
	receiverName = receiverId.Name
	receiverSocket = receiverId.ServerIp + ":" + receiverId.FileReceptionPort
	return
}

// AskForFilePath asks the user to enter the path to the file to be sent
func AskForFilePath() string {
	fmt.Printf("Please enter the path (must not contain any spaces) to the file you want to send:\n> ")
	var filePath string
	_, err := fmt.Scan(&filePath)
	if err != nil {
		log.Fatalf("Error scanning input: %s\n", err)
	}
	return filePath
}

func AskForSocket() string {
	fmt.Printf("Please enter the socket (<ip>:<port>) you want to send the file to:\n> ")
	var socket string
	_, err := fmt.Scan(&socket)
	if err != nil {
		log.Fatalf("Error scanning input: %s\n", err)
	}

	if !myutils.ValidateSocket(socket) {
		log.Fatalf("Invalid socket: %s\n", socket)
	}
	return socket
}
