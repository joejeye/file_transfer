package main

import (
	"encoding/json"
	"file_transfer_naive/formatting"
	filetransfer "file_transfer_naive/global_config"
	"file_transfer_naive/myutils"
	"fmt"
	"net"
)

func RespondToPeerDiscovery() {
	lockedConfig := filetransfer.GetLockedConfig()
	listenPort := lockedConfig.PDServerListenPort

	conn, err := net.ListenPacket("udp", ":"+listenPort)
	if err != nil {
		fmt.Printf("Error listening on port %s: %s\n", listenPort, err)
		return
	}
	defer func(conn net.PacketConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(conn)

	buffer := make([]byte, lockedConfig.PeerDiscoveryPacketSizeLimitBytes)
	fmt.Println("Listening for peer discovery messages on port", listenPort)
	myName := myutils.RandName()
	myutils.PrintMyInfo(myName)
	for {
		n, addr, err := conn.ReadFrom(buffer)
		msg := string(buffer[:n])
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			continue
		}

		fmt.Printf("Received message from %s: %s\n", addr.String(), msg)

		// Respond to the peer
		serverConfig := GetServerConfig()
		respMsg := formatting.PeerRespMsg{
			Name:              myName,
			FileReceptionPort: serverConfig.ListenPort,
		}
		respJSON, err := json.Marshal(respMsg)
		if err != nil {
			fmt.Println("Error marshalling response message:", err)
			continue
		}
		senderIP, _, err := net.SplitHostPort(addr.String())
		if err != nil {
			fmt.Printf("Error splitting host and port from address %s: %s\n", addr.String(), err)
			continue
		}
		respConn, err := net.Dial("udp", senderIP+":"+lockedConfig.PDClientListenPort)
		_, err = respConn.Write(respJSON)
		if err != nil {
			fmt.Printf("Error responding to %s:%s: %s\n", senderIP, lockedConfig.PDClientListenPort, err)
			continue
		} else {
			fmt.Printf("Responded to %s:%s with message: %s\n", senderIP, lockedConfig.PDClientListenPort, respJSON)
		}
	}
}
