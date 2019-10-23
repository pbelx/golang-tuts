package main

//tls receiving client non stop test
import "crypto/tls"

func main() {
	println("trying ssl client")
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:2222", conf)
	if err != nil {
		println("fail")
	}
	println("connected")
	// buf, err := conn.Write([]byte("Yes we on"))
	// if err != nil {
	// 	println("cant write")
	// }

	for {
		buf := make([]byte, 500)
		conn.Read(buf)
		println(string(buf))
	}

}
