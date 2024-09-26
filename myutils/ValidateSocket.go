package myutils

import (
	"strconv"
	"strings"
)

// ValidateSocket Checks if the socket follows the format a.b.c.d:port,
// where a, b, c, d are integers in the range [0, 255] and port is an integer in the range [0, 65535]
func ValidateSocket(socket string) bool {
	// Split the socket into IP and port
	socketParts := strings.Split(socket, ":")
	if len(socketParts) != 2 {
		return false
	}

	// Validate the IP
	ipParts := strings.Split(socketParts[0], ".")
	if len(ipParts) != 4 {
		return false
	}
	for _, part := range ipParts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}

	// Validate the port
	port, err := strconv.Atoi(socketParts[1])
	if err != nil || port < 0 || port > 65535 {
		return false
	}

	return true
}
