package myutils

import "log"

func MyPathJoin(paths ...string) string {
	result := ""
	myOS := GetOS()
	delim := ""
	if myOS == "windows" {
		delim = "\\"
	} else if myOS == "linux" {
		delim = "/"
	} else {
		log.Fatalf("Unsupported OS: %s. The app currently supports Windows and Linux.", myOS)
	}
	for i, path := range paths {
		if i == 0 {
			result += path
		} else {
			result += delim + path
		}
	}

	return result
}
