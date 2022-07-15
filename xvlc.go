package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	vlc "github.com/adrg/libvlc-go/v3"
)

func main() {

	fmt.Println("Music Steamer client Starting")
	url := os.Args[1]
	vlc.Init()
	player, err := vlc.NewPlayer()

	//fmt.Println("inits")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("inits")

	defer func() {
		player.Stop()
		player.Release()
	}()

	// Add a media file from path or from URL.
	// Set player media from path:
	// media, err := player.LoadMediaFromPath("localpath/test.mp4")
	// Set player media from URL:
	media, err := player.LoadMediaFromURL(url)
//	fmt.Println("inits2")

	if err != nil {
		//log.Fatal(err)
                fmt.Println("cannot open file")
	}
	defer media.Release()

	// Retrieve player event manager.
	manager, err := player.EventManager()
	if err != nil {
		log.Fatal(err)
	}

	// Register the media end reached event with the event manager.
	quit := make(chan struct{})
	eventCallback := func(event vlc.Event, userData interface{}) {
		close(quit)
	}

	eventID, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Detach(eventID)

	// Start playing the media.
	// if err = player.Play(); err != nil {
	// 	log.Fatal(err)
	// }
	player.Play()
	<-quit
}
