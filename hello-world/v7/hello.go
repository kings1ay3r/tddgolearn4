package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name

}

func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case "es":
		prefix = spanishHelloPrefix
	case "fr":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("World", ""))
}
