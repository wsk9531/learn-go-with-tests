package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const teReo = "Te Reo"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const teReoHelloPrefix = "Kia Ora, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case teReo:
		prefix = teReoHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("mate", spanish))
}
