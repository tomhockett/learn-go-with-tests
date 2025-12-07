package main


var languagePrefixes = map[string]string{
    "English": "Hello, ",
    "Spanish": "Hola, ",
    "French":  "Bonjour, ",
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) string {
    if prefix, ok := languagePrefixes[language]; ok {
        return prefix
    }
    return languagePrefixes["English"]
}
