package main

//Fool google devices into thinking they are online
import (
	"fmt"
	"net/http"
)

func googleFool(w http.ResponseWriter, r *http.Request) {

	fmt.Println("got hit")

	w.WriteHeader(http.StatusNoContent)
}

func main() {

	http.HandleFunc("/generate_204", googleFool)
	fmt.Println("server up port 80")
	http.ListenAndServe(":80", nil)
}
