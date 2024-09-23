package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: main server_ip:server_port path_to_file")
		return
	}

	serverAddr := os.Args[1]
	filePath := os.Args[2]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Connect to the server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

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
