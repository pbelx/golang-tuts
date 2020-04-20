package main

import (
	"fmt"
	"net"
)

func main() {
	for port := 1; port < 1024; port++ {
		address := fmt.Sprintf("localhost:%d", port)
		// fmt.Println(address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// fmt.Println("port closed")
			continue
		}
		conn.Close()
		fmt.Printf("port open %d\n", port)
	}
}
