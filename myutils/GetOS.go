package myutils

import (
	"log"
	"runtime"
	"strings"
)

func GetOS() string {
	myOS := runtime.GOOS
	if strings.Contains(myOS, "linux") {
		return "linux"
	} else if strings.Contains(myOS, "windows") {
		return "windows"
	} else {
		log.Fatalf("Unsupported OS: %s. The app currently supports Windows and Linux.", myOS)
		return "" // This line will never be reached
	}
}
