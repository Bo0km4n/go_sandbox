package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:2345")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", ":0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)

		buf = make([]byte, 1024)
		n, addr, err := Conn.ReadFrom(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
	}
}
