package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	englishHello       = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// const englishHello = "Hello, "
// const spanish = "Spanish"
// const spanishHelloPrefix = "Hola, "
// const french = "French"
// const frenchHelloPrefix = "Bonjour, "

func Hello1(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHello + name
}

// func Hello(name, language string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	prefix := englishHello
// 	switch language {
// 	case spanish:
// 		prefix = spanishHelloPrefix
// 	case french:
// 		prefix = frenchHelloPrefix

// 	}
// 	return prefix + name

// }

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingsPrefix(language) + name

}

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHello
	}
	return

}

func main() {
	fmt.Println(Hello("world", ""))

}
