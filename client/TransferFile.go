package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func TransferFile(receiverSocket string, filePath string) {
	// Open the file
	if strings.Contains(filePath, " ") {
		log.Fatalln("Invalid input: file path must not contain any spaces")
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	// Connect to the server
	conn, err := net.Dial("tcp", receiverSocket)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(conn)

	// Send the filename
	_, err = conn.Write([]byte(filepath.Base(filePath) + "\n"))
	if err != nil {
		fmt.Println("Error sending filename:", err)
		return
	}

	// Send the file content
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file:", err)
		return
	}

	// Read the server's response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading server response:", err)
		return
	}

	// Print the server's response
	fmt.Println(string(buffer[:n]))
}
