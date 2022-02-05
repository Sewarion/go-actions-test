package main

import (
	"fmt"

	"github.com/Sewarion/go-actions-test/greet"
)

func main() {
	name := "Julian"
	fmt.Println(greet.Greeting(name))
}
