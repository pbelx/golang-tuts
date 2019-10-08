package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/proxy"
)

func main() {
	dialsocks, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		fmt.Println("got error on proxy")
	}

	tr := &http.Transport{Dial: dialsocks.Dial}
	myclient := &http.Client{
		Transport: tr,
	}
	resp, err := myclient.Get("http://icanhazip.com")
	if err != nil {
		fmt.Println("failed to fe")
	}

	fmt.Println(resp)
}
