package greet

import "fmt"

func Greeting(name string) string {
	greet := fmt.Sprintf("Hello, %s!", name)
	return greet
}
