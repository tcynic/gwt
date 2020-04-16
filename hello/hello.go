package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello defines a string that it returns
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
