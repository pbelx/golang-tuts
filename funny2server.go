package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	socket, _ := net.Listen("tcp", ":9000")
	socket2, _ := net.Listen("tcp", ":9001")
	fmt.Println("server up on 9000")
	for {
		conn, _ := socket.Accept()
		conn.Write([]byte("belski server\n"))
		conn2, _ := socket2.Accept()
		conn2.Write([]byte("belski 222 server\n"))
		go reply(conn, conn2)
	}

}

func reply(konn net.Conn, kon2 net.Conn) {
	go io.Copy(kon2, konn)
	go io.Copy(konn, kon2)
}
