package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"h12.io/socks"
)

func main() {
	dialsocks := socks.Dial("socks5://127.0.0.1:1080?timeout=10s")
	tr := &http.Transport{Dial: dialsocks}
	httpclient := &http.Client{Transport: tr}
	fmt.Println("connecting to site")
	resp, err := httpclient.Get(os.Args[1])

	if err != nil {
		fmt.Println("got erro codes")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("got no response")
	}
	fmt.Println(string(body))
}
