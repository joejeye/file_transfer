package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func main() {
	dlDir := getDownloadDir()
	fmt.Println("Download directory:", dlDir)
	listenPort := getListenPort()
	fmt.Println("Listen port:", listenPort)

	listener, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on: " + listenPort)

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
	defer conn.Close()

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
		conn.Write([]byte("Server error: File exists already"))
		return
	}
	respMsg := fmt.Sprintf(
		"Server is downloading %s to %s. Please check the server-side output"+
			" for the confirmation of file reception.",
		filename, downloadPath)
	conn.Write([]byte(respMsg))

	// Create the file
	file, err := os.Create(downloadPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		conn.Write([]byte("Error creating file"))
		return
	}
	defer file.Close()

	// Save the file content
	_, err = io.Copy(file, reader)
	if err != nil {
		fmt.Println("Error saving file:", err)
		conn.Write([]byte("Error saving file"))
		return
	}

	fmt.Printf("File %s downloaded to %s\n", filename, downloadPath)
}

// Parse the configuration file
type Config struct {
	DlDir      string `toml:"download_directory"`
	ListenPort string `toml:"listen_port"`
}

func getDownloadDir() string {
	var conf Config
	_, err := toml.DecodeFile("server/config.toml", &conf)
	if err != nil {
		panic(fmt.Errorf("error decoding config file: %w", err))
	}
	return conf.DlDir
}

func getListenPort() string {
	var conf Config
	_, err := toml.DecodeFile("server/config.toml", &conf)
	if err != nil {
		panic(fmt.Errorf("error decoding config file: %w", err))
	}
	return conf.ListenPort
}
