package main

import (
	"encoding/json"
	"file_transfer_naive/formatting"
	filetransfer "file_transfer_naive/global_config"
	"file_transfer_naive/myutils"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

const (
	BroadcastAddr = "255.255.255.255"
)

func PeerDiscovery() []formatting.ServerID {
	myName := myutils.RandName()
	myutils.PrintMyInfo(myName)

	// Create UDP connection for broadcasting
	lockedConfig := filetransfer.GetLockedConfig()
	broadcastConn, err := net.Dial("udp", BroadcastAddr+":"+lockedConfig.PDServerListenPort)
	if err != nil {
		fmt.Println("Error dialing broadcast address:", err)
		return nil
	}
	defer func() {
		err := broadcastConn.Close()
		if err != nil {
			fmt.Println("Error closing broadcast connection:", err)
		}
	}()

	// Create UDP connection for listening responses to the broadcast
	listenAddr, err := net.ResolveUDPAddr("udp", ":"+lockedConfig.PDClientListenPort)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return nil
	}
	listenConn, err := net.ListenPacket("udp", listenAddr.String())
	if err != nil {
		fmt.Println("Error listening on port", lockedConfig.PDClientListenPort, ":", err)
		return nil
	}
	defer func() {
		err := listenConn.Close()
		if err != nil {
			fmt.Println("Error closing listen connection:", err)
		}
	}()

	// Start goroutine to listen for responses
	var wg sync.WaitGroup
	wg.Add(1)
	timedOut := false
	go func() { // The client waits for a certain amount of time before closing the connection
		defer wg.Done()
		time.Sleep(time.Duration(lockedConfig.PDTimeOutMilliSec) * time.Millisecond)
		timedOut = true
	}()
	var serverIds []formatting.ServerID
	go func() {
		buffer := make([]byte, lockedConfig.PeerDiscoveryPacketSizeLimitBytes)
		for {
			if timedOut {
				return
			}
			n, addr, err := listenConn.ReadFrom(buffer)
			JsonMsg := buffer[:n]
			if err != nil {
				if timedOut { // This seemingly redundant check is necessary to prevent the program from panicking
					return
				}
				log.Println("Error reading from connection:", err)
				continue
			}
			respMsg := formatting.PeerRespMsg{}
			err = json.Unmarshal(JsonMsg, &respMsg)
			if err != nil {
				fmt.Printf("Error decoding JSON message from %s: %s\n", addr.String(), err)
				continue
			}
			serverIp, _, err := net.SplitHostPort(addr.String())
			serverId := formatting.ServerID{
				PeerRespMsg: respMsg,
				ServerIp:    serverIp,
			}
			serverIds = append(serverIds, serverId)
		}
	}()

	// Broadcast the peer discovery message
	broadcastMsg := formatting.PeerRespMsg{
		Name:              myName,
		FileReceptionPort: "",
	}
	JsonMsg, err := json.Marshal(broadcastMsg)
	if err != nil {
		log.Println("Error marshalling broadcast message:", err)
		return nil
	}
	_, err = broadcastConn.Write(JsonMsg)
	if err != nil {
		fmt.Println("Error writing to broadcast connection:", err)
		return nil
	} else {
		fmt.Printf("Broadcasted message to %s:%s\n", BroadcastAddr, lockedConfig.PDServerListenPort)
	}

	wg.Wait()
	return serverIds
}
