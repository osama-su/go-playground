package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
// Note: the function name begins with an uppercase letter, making it accessible to code outside the package.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling the Hello function to get a message for each name.
	// Assign the message to the map, using the name as the key.
	// Note: The range keyword returns two values: the index of the current item in the loop and a copy of the item's value.
	// You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random.
// Note: the function name begins with a lowercase letter, making it accessible only to code in its own package.
func randomFormat() string {
	// a slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
