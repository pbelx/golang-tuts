package main

import "fmt"

//Car info
type Car struct {
	Driver string
	Plate  int
}

//Road type
type Road struct{}

func (r *Road) hoot() {
	fmt.Println("You drive on tarmac!! covid")
}

func (c *Car) hoot() {
	fmt.Println("driver is ", c.Driver, "plate is ", c.Plate)
}

//Friend interface
type Friend interface {
	hoot()
}

//Greet function
func Greet(f Friend) {
	f.hoot()
}

func main() {
	var cc = new(Car)
	cc.Driver = "Ken"
	cc.Plate = 9112
	Greet(cc)
	var rr = new(Road)
	Greet(rr)
}
