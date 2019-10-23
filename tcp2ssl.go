package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:5555")
	errorcheck(err, "cannot listen to port")
	fmt.Println("server on 5555")
	for {
		socket, err := server.Accept()
		errorcheck(err, "cannot accept")
		fmt.Println("got connection")

		go consend(socket)

	}
}

func consend(socket net.Conn) {
	defer socket.Close()

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	remote, err := tls.Dial("tcp", "127.0.0.1:2222", conf)
	if err != nil {
		println("fail")
	}
	println("connected")

	io.Copy(remote, socket)
	io.Copy(socket, remote)

}

func errorcheck(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}

}
