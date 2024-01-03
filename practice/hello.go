package main

const EnglishPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return EnglishPrefix + name
}
