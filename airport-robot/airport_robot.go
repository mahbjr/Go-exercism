package airportrobot

import "fmt"

// Greeter interface allows the robot to greet in different languages
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

// SayHello makes the robot greet a visitor in their native language
func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

// Italian implements the Greeter interface for Italian language
type Italian struct{}

// LanguageName returns the name of the language
func (i Italian) LanguageName() string {
	return "Italian"
}

// Greet returns a greeting in Italian
func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// Portuguese implements the Greeter interface for Portuguese language
type Portuguese struct{}

// LanguageName returns the name of the language
func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

// Greet returns a greeting in Portuguese
func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
