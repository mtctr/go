package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return fmt.Sprint(englishHelloPrefix, "World")
	}

	return fmt.Sprint(englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("World"))
}
