package main

//Check if music streams are active
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// func checkStream(r *http.Response) {
// 	if r.Header.Get("Content-Type") == "application/octet-stream" {
// 		//stream is playing
// 		fmt.Println("Stream is active")

// 	} else {
// 		//stream is offline
// 		fmt.Println("stream is down")
// 	}
// }
func surfStream(url string) {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return

	}
	if r.Header.Get("Content-Type") == "audio/mpeg" || r.Header.Get("Content-Type") == "audio/aacp" {
		//stream is playing
		fmt.Println("Stream is active " + url)
		fmt.Println(r.Header.Get("Content-Type"))

	} else {
		//stream is offline
		fmt.Println("stream is down " + url)
	}
}

type Rstations struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Id   string `json:"id"`
}

func main() {
	fmt.Println("music client")
	// resp, err := http.Get("http://techwiz.icu:8080/rxfm")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println(resp.Header)
	// 	return
	// }
	fileSpecial, _ := os.ReadFile("./radiostations.json")
	var stations []Rstations
	json.Unmarshal(fileSpecial, &stations)

	for _, v := range stations {

		fmt.Println(v.Url)
		surfStream(v.Url)

	}

	// checkStream(resp)

}
