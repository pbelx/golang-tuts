package main

//GET news headlines
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"media,attr"`
	Channel struct {
		Text          string `xml:",chardata"`
		Generator     string `xml:"generator"`
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Language      string `xml:"language"`
		WebMaster     string `xml:"webMaster"`
		Copyright     string `xml:"copyright"`
		LastBuildDate string `xml:"lastBuildDate"`
		Description   string `xml:"description"`
		Item          []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
			Guid  struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate     string `xml:"pubDate"`
			Description string `xml:"description"`
			Source      struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
			} `xml:"source"`
		} `xml:"item"`
	} `xml:"channel"`
}

func main() {
	fmt.Println("New Client Starting...")

	resp, err := http.Get("https://news.google.com/rss/search?q=news&hl=en-UG&gl=UG&ceid=UG:en")
	if err != nil {
		fmt.Println("got Err")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err")
	}
	var news Rss

	xml.Unmarshal(data, &news)
	for _, v := range news.Channel.Item {
		fmt.Println(v.Title)
		fmt.Println(v.Link)

	}

	// fmt.Println(news.Channel.Item)

}
