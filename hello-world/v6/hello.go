package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == "es" {
		return spanishHelloPrefix + name
	}
	if lang == "fr" {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("World", ""))
}
