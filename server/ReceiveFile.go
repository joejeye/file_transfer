package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func ReceiveFile() {
	dlDir := getDownloadDir()
	fmt.Println("Download directory:", dlDir)
	listenPort := getListenPort()

	listener, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Error closing listener:", err)
		}
	}(listener)

	fmt.Println("Listening for receiving files on port", listenPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, dlDir)
	}
}

func handleConnection(conn net.Conn, dlDir string) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Read the filename
	reader := bufio.NewReader(conn)
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading filename:", err)
		return
	}
	filename = filename[:len(filename)-1] // Remove newline character

	// Check if file exists
	downloadPath := filepath.Join(dlDir, filename)
	if _, err := os.Stat(downloadPath); err == nil {
		_, err := conn.Write([]byte("Server error: File exists already"))
		if err != nil {
			return
		}
		return
	}
	respMsg := fmt.Sprintf(
		"Server is downloading %s to %s. Please check the server-side output"+
			" to confirm the file reception.",
		filename, downloadPath)
	_, err = conn.Write([]byte(respMsg))
	if err != nil {
		return
	}

	// Create the file
	file, err := os.Create(downloadPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		_, err := conn.Write([]byte("Error creating file"))
		if err != nil {
			return
		}
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file:", err)
		}
	}(file)

	// Save the file content
	_, err = io.Copy(file, reader)
	if err != nil {
		fmt.Println("Error saving file:", err)
		_, err := conn.Write([]byte("Error saving file"))
		if err != nil {
			return
		}
		return
	}

	fmt.Printf("File %s downloaded to %s\n", filename, downloadPath)
}

func getDownloadDir() string {
	conf := GetServerConfig()
	return conf.DlDir
}

func getListenPort() string {
	conf := GetServerConfig()
	return conf.ListenPort
}
