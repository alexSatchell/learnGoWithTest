package main

import "fmt"

// It is good to seperate domain code from the outside World
// named returns
// in go, functinos names starting with a capital letter are public
// && the opposite is private functions
const (
	spanish            = "Spanish"
	french             = "French"
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("", ""))
}
