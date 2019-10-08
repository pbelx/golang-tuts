package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/proxy"
)

func main() {
	dialsocks, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		fmt.Println("got error on proxy")
	}
	// fmt
	tr := &http.Transport{Dial: dialsocks.Dial}
	myclient := &http.Client{
		Transport: tr,
	}
	resp, err := myclient.Get("http://goo.gl")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read body")
	}
	fmt.Println(string(body))
}
