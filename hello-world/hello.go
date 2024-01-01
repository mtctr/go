package main

import "fmt"

func Hello(name string) string {
	return fmt.Sprint("Hello, ", name)
}

func main() {
	fmt.Println(Hello("world"))
}
