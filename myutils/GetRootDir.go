package myutils

import (
	"path/filepath"
	"runtime"
)

func GetRootDir() string {
	_, caller0, _, _ := runtime.Caller(0) // <root>/myutils/GetRootDir.go
	myutilsDir := filepath.Dir(caller0)   // <root>/myutils
	rootDir := filepath.Dir(myutilsDir)   // <root>
	return rootDir
}
