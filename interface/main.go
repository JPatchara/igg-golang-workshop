package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}
type spanishBot struct {
}
type thaiBot struct {
}

func main() {
	name := "Jin"
	eb := englishBot{}
	sb := spanishBot{}
	tb := thaiBot{}

	printGreeting(eb, name)
	printGreeting(sb, name)
	printGreeting(tb, name)
	// printEbGreeting(eb)
	// printSbGreeting(sb)
	// printTbGreeting(tb)
}

func printGreeting(b bot, name string) {
	fmt.Println(b.getGreeting(), ",", name)
}

func (eb englishBot) getGreeting() string {
	return "Hello! Bro."
}
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
func (tb thaiBot) getGreeting() string {
	return "Swatdee jaa!"
}

// func printEbGreeting(eb englishBot) {
// 	fmt.Println(eb.getEbGreeting())
// }
// func printSbGreeting(sb spanishBot) {
// 	fmt.Println(sb.getSbGreeting())
// }
// func printTbGreeting(tb thaiBot) {
// 	fmt.Println(tb.getTbGreeting())
// }
