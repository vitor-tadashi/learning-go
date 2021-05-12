package main

const englishHelloPrefix = "Hello, "
const spanisHelloPrefix = "Hola, "
const portugueseHelloPrefix = "Olá, "

const spanish = "spanish"
const portuguese = "portuguese"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanisHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
