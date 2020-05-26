package main

import (
	"fmt"
	t "time"
)

const prefix = "Hello, "

//Hello function
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Baby"))
	fmt.Println(t.Now())
}
