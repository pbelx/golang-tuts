package main
//play mp3 stream given url args
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

func run() error {

	

	fmt.Println("Music server starting")


	musicurl := os.Args[1]
	resp, _ := http.Get(musicurl)

	d, err := mp3.NewDecoder(resp.Body)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	fmt.Printf("Length: %d[bytes]\n", d.Length())

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
