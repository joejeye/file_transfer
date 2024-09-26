package main

import (
	"fmt"
	"testing"
)

func TestGetServerConfig(t *testing.T) {
	sConfig := GetServerConfig()
	fmt.Printf("Server config: %+v\n", sConfig)
}
