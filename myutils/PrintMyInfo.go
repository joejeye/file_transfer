package myutils

import "fmt"

func PrintMyInfo(myName string) {
	myIp := GetMyIp()
	fmt.Printf("I am %s (%s)\n", myName, myIp.String())
}
