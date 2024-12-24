package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
)

func main() {
    port := flag.Int("port", 8080, "port to serve on")
    dir := flag.String("dir", ".", "directory to serve")
    flag.Parse()

    fs := http.FileServer(http.Dir(*dir))
    http.Handle("/", fs)

    addr := fmt.Sprintf(":%d", *port)
    log.Printf("Serving %s on HTTP port: %d\n", *dir, *port)
    log.Fatal(http.ListenAndServe(addr, nil))
}
