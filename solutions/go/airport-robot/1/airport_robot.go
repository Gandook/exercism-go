package airportrobot

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(string) string
}

func SayHello(name string, greeter Greeter) string {
    return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}

type Italian struct {
}

func (italian Italian) LanguageName() string {
    return "Italian"
}

func (italian Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {
}

func (portuguese Portuguese) LanguageName() string {
    return "Portuguese"
}

func (portuguese Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name)
}
