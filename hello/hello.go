package main

import (
	"fmt"
	t "time"
)

// Hello function
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Baby"))
	fmt.Println(t.Now())
}
