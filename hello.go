package main

import (
	"fmt"
	"strings"
)

// const englishHelloPrefix = "Hello"

// func Hello() string {
// 	return "Hello, world!"
// }

func HelloWithName(name string, lang string) string {

	if name == "" {
		name = "world"
	}

	if strings.EqualFold(name, "fernanda") {
		name = "world"
	}

	return GetLanguagePrefix(lang) + ", " + name
}

func GetLanguagePrefix(lang string) (prefix string) {

	var defaultHelloPrefix string
	switch lang {
	case "Spanish":
		defaultHelloPrefix = "Hola"
	case "English":
		defaultHelloPrefix = "Hello"
	case "French":
		defaultHelloPrefix = "Bonjour"
	}

	return defaultHelloPrefix

}

func main() {
	// fmt.Println(Hello("Carlos"))
	// fmt.Println(Hello())
	fmt.Println(HelloWithName("fernanda", "Spanish"))
	fmt.Println(HelloWithName("fernanda", "English"))
}
