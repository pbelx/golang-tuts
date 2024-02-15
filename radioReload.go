package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Station struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var stations []Station

func loadStations(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &stations)
	if err != nil {
		return err
	}

	return nil
}

func handleStation(c *gin.Context) {
	stationName := c.Param("stationName")
	for _, station := range stations {
		if station.Name == stationName {
			resp, err := http.Get(station.URL)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer resp.Body.Close()
			io.Copy(c.Writer, resp.Body)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Station not found"})
}

func handleAllStations(c *gin.Context) {
	var names []string
	for _, station := range stations {
		names = append(names, station.Name)
	}
	c.JSON(http.StatusOK, gin.H{"stations": names})
}

func reloadStationsPeriodically(filename string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := loadStations(filename)
			if err != nil {
				fmt.Println("Error reloading stations:", err)
			}
		}
	}
}

func main() {
	err := loadStations("stations.json")
	if err != nil {
		fmt.Println("Error loading stations:", err)
		return
	}

	go reloadStationsPeriodically("stations.json", 5*time.Second)

	fmt.Println("gin go")
	r := gin.Default()
	r.GET("/station/:stationName", handleStation)
	r.GET("/all", handleAllStations)

	fmt.Println("Server up port 8080")
	r.Run()
}
