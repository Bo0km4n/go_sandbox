package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("メッセージを入力してください")
	var message string
	fmt.Scan(&message)

	var server_ip = "127.0.0.1"
	var server_port = "3333"
	var client_ip = "127.0.0.1"
	var client_port = 33334

	tcp_addr, err := net.ResolveTCPAddr("tcp", server_ip+":"+server_port)
	checkError(err)
	client_addr := new(net.TCPAddr)
	client_addr.IP = net.ParseIP(client_ip)
	client_addr.Port = client_port

	conn, err := net.DialTCP("tcp", client_addr, tcp_addr)
	checkError(err)
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))

	readBuf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readlen, err := conn.Read(readBuf)
	checkError(err)

	fmt.Println("server: " + string(readBuf[readlen]))

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error %s", err.Error())
		os.Exit(1)
	}
}
