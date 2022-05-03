package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	value, _ := ioutil.ReadAll(body)
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func sanyu(c *gin.Context) {
	resp, _ := http.Get("https://techwiz.icu:8443/sanyu")
	io.Copy(c.Writer, resp.Body)
}
func hot100(c *gin.Context) {
	resp, _ := http.Get("https://techwiz.icu:8443/hot100")
	io.Copy(c.Writer, resp.Body)
}
func rxfm(c *gin.Context) {
	resp, _ := http.Get("https://techwiz.icu:8443/rxfm")
	io.Copy(c.Writer, resp.Body)
}
func hii(c *gin.Context) {
	resp, _ := http.Get("https://techwiz.icu:8443/hii")
	io.Copy(c.Writer, resp.Body)
}
func rcity(c *gin.Context) {
	resp, _ := http.Get("https://techwiz.icu:8443/rcity")
	io.Copy(c.Writer, resp.Body)
}
func ugdjs(c *gin.Context) {
	resp, _ := http.Get("https://techwiz.icu:8443/ugdjs")
	io.Copy(c.Writer, resp.Body)
}

func main() {
	fmt.Println("gin go")
	r := gin.Default()
	r.GET("/", test)
	r.POST("/login", logindetails)
	r.POST("/bod", bodyfn)
	r.GET("/sanyu", sanyu)
	r.GET("/hot100", hot100)
	r.GET("/rxfm", rxfm)
	r.GET("/hii", hii)
	r.GET("/rcity", rcity)
	r.GET("/ugdjs", ugdjs)

	fmt.Println("Server up port 8080")
	r.Run()
}
