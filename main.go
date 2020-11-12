package main

import (
	"fmt"
	"github.com/owendland/learn_go_with_tests/di"
	"github.com/owendland/learn_go_with_tests/mocking"
	"net/http"
	"os"
	"time"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	sleeper := mocking.NewConfigurableSleeper(1*time.Second, time.Sleep)
	err := mocking.Countdown(os.Stdout, sleeper)
	if err != nil {
		fmt.Println("Countdown failed: " + err.Error())
	}

	err = http.ListenAndServe(":5000", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := di.Greet(w, "world")
			if err != nil {
				panic(nil)
			}
		}))
	if err != nil {
		fmt.Println("Server error occurred: " + err.Error())
	}
}
