package main

import (
  
  "net/http"
  "fmt"
)

func main() {
  fs := http.FileServer(http.Dir("."))
  http.Handle("/", fs)

  fmt.Println("Listening...8000")
  http.ListenAndServe(":8000", nil)
}
