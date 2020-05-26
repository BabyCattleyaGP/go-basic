package main

import (
	"fmt"
	t "time"
)

const prefix = "Hello, "
const frenchPrefix = "Bonjour, "

//Hello function
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "French" {
		return frenchPrefix + name
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Baby", "French"))
	fmt.Println(t.Now())
}
