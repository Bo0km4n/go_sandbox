package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	port := ":3333"
	tcp_addr, err := net.ResolveTCPAddr("tcp", port)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcp_addr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	fmt.Println("client accept!")
	messageBuf := make([]byte, 1024)
	mes_len, err := conn.Read(messageBuf)
	checkError(err)

	message := string(messageBuf[:mes_len])
	fmt.Println("client message: " + message)
	message = message + "too!"

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))
}

func checkError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
