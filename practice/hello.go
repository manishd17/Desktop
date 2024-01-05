package main

const (
	Spanish       = "spanish"
	EnglishPrefix = "Hello, "
	Spanishprefix = "Hola, "
	FrenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	//case "French":
	//	FrenchPrefix
	//case "Spanish":
	//	Spanishprefix
	//default:
	//	EnglishPrefix
	}

	return
}
