package main

import (
	"fmt"
	t "time"
)

const englishPrefix = "Hello, "
const frenchPrefix = "Bonjour, "
const spanishPrefix = "Hola, "

//Hello function
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchPrefix
	case "spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Baby", "french"))
	fmt.Println(t.Now())
}
