package main

import (
	"fmt"
	"net"
)

func main() {

	server, err := net.Listen("tcp", "127.0.0.1:5555")
	checkerror(err, "failed to listen")
	print("server up on 5555")

	for {
		connection, err := server.Accept()
		checkerror(err, "failed to accept connection")
		print("accepting connections now")

		go handleclient(connection)

	}

}

func handleclient(connection net.Conn) {
	defer connection.Close()
	buf := make([]byte, 1024)
	connection.Read(buf)
	fmt.Println(string(buf))
	connection.Write([]byte(string("got it\n")))
	fmt.Println("sent message")
	// connection.Close()
}

func checkerror(err error, msg string) {
	if err != nil {
		fmt.Println(string(msg))
	}

}

func print(msg string) {
	fmt.Println(msg)
}
