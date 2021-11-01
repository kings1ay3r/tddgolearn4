package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	switch lang {
	case "es":
		return spanishHelloPrefix + name
	case "fr":
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func greetingPrefix()
func main() {
	fmt.Println(Hello("World", ""))
}
