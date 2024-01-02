package main

import "fmt"

const english = "English"
const englishHelloPrefix = "Hello, "

const portuguese = "Portuguese"
const portugueseHelloPrefix = "Olá, "

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {

	prefix := getPrefix(language)

	if name == "" {
		return fmt.Sprint(prefix, "World")
	}

	return fmt.Sprint(prefix, name)
}

func getPrefix(language string) string {
	switch language {
	case portuguese:
		return portugueseHelloPrefix
	case spanish:
		return spanishHelloPrefix
	default:
		return englishHelloPrefix
	}

}
func main() {
	fmt.Println(Hello("World", "English"))
}
