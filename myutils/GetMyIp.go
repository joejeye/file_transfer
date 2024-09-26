package myutils

import "net"

func GetMyIp() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet:
			if v.IP.IsGlobalUnicast() {
				return v.IP
			}
		}
	}
	panic("No global unicast IP address found")
}
