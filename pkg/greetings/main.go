package greetings

import "fmt"

func Get(name string) string {
	message := fmt.Sprintf("Hi, %v", name)
	return message
}
