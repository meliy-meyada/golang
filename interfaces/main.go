package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type thaiBot struct{}

func main() {
	eb := englishBot{}
	tb := thaiBot{}


	printGreeting(eb)
	printGreeting(tb)

}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}


func (englishBot) getGreeting() string {
	// Very Custom logic for generating an english greeting

	return "Hi there!"

}
func (thaiBot) getGreeting() string {

	return "สวัสดี!"
}