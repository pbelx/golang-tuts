package main

//Check if music streams are active
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Rstatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

var rs []Rstatus

func surfStream(url, name string) {

	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return

	}
	if r.Header.Get("Content-Type") == "audio/mpeg" || r.Header.Get("Content-Type") == "audio/aacp" {
		//stream is playing
		fmt.Println("Stream is active " + url)
		rs = append(rs, Rstatus{Name: name, Status: "active"})
		// fmt.Println(r.Header.Get("Content-Type"))

	} else {
		//stream is offline
		rs = append(rs, Rstatus{Name: name, Status: "offline"})

		fmt.Println("stream is down " + url)
	}
	bytes, _ := json.Marshal(rs)
	_ = ioutil.WriteFile("/tmp/status.json", bytes, 0644)

}

type Rstations struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Id   string `json:"id"`
}

func main() {
	fmt.Println("music client")
	ticker := time.NewTicker(30 * time.Minute)

	for t := range ticker.C {
		rs = nil
		fmt.Println("Invoked at ***************\n", t)
		fileSpecial, _ := os.ReadFile("./services.json")
		var stations []Rstations
		json.Unmarshal(fileSpecial, &stations)

		for _, v := range stations {

			fmt.Println(v.Url)
			surfStream(v.Url, v.Name)

		}
	}

}
