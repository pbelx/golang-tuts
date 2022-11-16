package main

import (
	// "log"

	// "encoding/json"
	"fmt"

	// ws "redirect/whatsapp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Homefunction(c *gin.Context) {
	c.Redirect(301, "https://techwiz.icu")
	path := c.Param("id")
	// agent := c.Request.UserAgent()
	// ip := c.RemoteIP()
	fmt.Println(path)
	// reqy := Requestinfo{Url: path, Useragent: agent, Ip: ip}
	// v, _ := json.Marshal(reqy)
	// ws.SendMessage(string(v))
	// fmt.Println(reqy)

}
func Homefunction2(c *gin.Context) {
	c.Redirect(301, "https://techwiz.icu")
	// path := c.Param("id")
	// agent := c.Request.UserAgent()
	// ip := c.RemoteIP()
	// fmt.Println(path)
	// reqy := Requestinfo{Useragent: agent, Ip: ip}
	// v, _ := json.Marshal(reqy)
	// ws.SendMessage(string(v))
	// fmt.Println(reqy)

}
func Demofunction(c *gin.Context) {
	c.Redirect(301, "http://techwiz.icu:8880")
	// path := c.Param("id")
	// agent := c.Request.UserAgent()
	// ip := c.RemoteIP()
	// fmt.Println(path)
	// reqy := Requestinfo{Useragent: agent, Ip: ip}
	// v, _ := json.Marshal(reqy)
	// ws.SendMessage(string(v))
	// fmt.Println(reqy)

}

// type Requestinfo struct {
// 	Url       string `json:"path"`
// 	Useragent string `json:"agent"`
// 	Ip        string `json:"ip"`
// }

func main() {

	fmt.Println("Redirect server up 80")
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/:id", Homefunction)
	r.GET("/", Homefunction2)
	r.GET("/demo", Demofunction)

	r.Run(":80")

}
