package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {

	for {
		print("kest")
		time.Sleep(3 * time.Second)
		file, _ := ioutil.ReadFile("/etc/passwd")
		fmt.Printf("%s", file)
		time.Sleep(2 * time.Second)
		fmt.Println("xfdf")
		host, _ := ioutil.ReadFile("/etc/hosts")
		time.Sleep(2 * time.Second)
		fmt.Printf("%s", host)
	}

}

func print(msg string) {
	fmt.Println(msg)
}
