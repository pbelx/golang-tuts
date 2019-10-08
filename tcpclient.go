package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Println("got error connecting")

	}
	fmt.Println("connected to host")

	conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	reply, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Print("got nothing")
	}
	fmt.Println(string(reply))
}
