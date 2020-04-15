package main

import (
	"fmt"
	t "time"
)

// Hello function
func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(t.Now())
}
