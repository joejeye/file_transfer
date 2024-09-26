package myutils

import "testing"

func TestGetMyIp(t *testing.T) {
	ip := GetMyIp()
	t.Logf("My IP: %s", ip)
}
