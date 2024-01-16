package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type Postdata struct {
	Url string
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func logindetails(c *gin.Context) {
	xname := c.Query("name")
	c.JSON(200, gin.H{
		"message": xname,
	})
}

func bodyfn(c *gin.Context) {
	body := c.Request.Body
	value, _ := io.ReadAll(body)
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func sanyu(c *gin.Context) {
	resp, _ := http.Get("https://techwiz.icu:8443/sanyu")
	io.Copy(c.Writer, resp.Body)
}
//post Url in json with youtube url link and have vlc play it 360p
func bodyFunction(c *gin.Context) {

	var requestBody Postdata

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Println("crazy err")
	}

	fmt.Println(requestBody.Url)
	c.JSON(200, gin.H{
		"message": string(requestBody.Url),
	})
	command := "yt-dlp"
	cmd := exec.Command(command, "-g", "-f 18 ", requestBody.Url)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	newurl := strings.ReplaceAll(string(output), "\n", "")
	vlcrun := exec.Command("vlc", string(newurl))

	vlcrun.Run()

}

func main() {
	fmt.Println("gin go")
	r := gin.Default()
	r.GET("/", test)
	r.GET("/sanyu", sanyu)
	r.POST("/body", bodyfn)
	r.POST("/url2", logindetails)
	r.POST("/url", bodyFunction)

	fmt.Println("Server up port 8080")
	r.Run()
}
