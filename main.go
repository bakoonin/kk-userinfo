package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"time"
)

func getPreferredIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

func main() {
	timeLayout := "20060102 15:04:05"

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now().Format(timeLayout)

	fmt.Println("Hostname:  " + getHostname())
	fmt.Println("IP:        " + getPreferredIP())
	fmt.Println("MAC:       " + getMacAddr())
	fmt.Println("Username:  " + user.Username)
	fmt.Println("Name:      " + user.Name)
	fmt.Println("UserID:    " + user.Uid)
	fmt.Println("HomeDir:   " + user.HomeDir)
	fmt.Println("LoginTime: " + now)
}
