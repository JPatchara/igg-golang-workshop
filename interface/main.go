package main

import (
	"fmt"
	"io"
	"net/http"
)

// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct {
// }
// type spanishBot struct {
// }
// type thaiBot struct {
// }

func main() {
	resp, err := http.Get("https://www.google.com")
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(resp.Header["Content-Type"])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
	// name := "Jin"
	// eb := englishBot{}
	// sb := spanishBot{}
	// tb := thaiBot{}

	// printGreeting(eb, name)
	// printGreeting(sb, name)
	// printGreeting(tb, name)
	// printEbGreeting(eb)
	// printSbGreeting(sb)
	// printTbGreeting(tb)
}

// func printGreeting(b bot, name string) {
// 	fmt.Println(b.getGreeting(), ",", name)
// }

// func (eb englishBot) getGreeting() string {
// 	return "Hello! Bro."
// }
// func (sb spanishBot) getGreeting() string {
// 	return "Hola!"
// }
// func (tb thaiBot) getGreeting() string {
// 	return "Swatdee jaa!"
// }

// func printEbGreeting(eb englishBot) {
// 	fmt.Println(eb.getEbGreeting())
// }
// func printSbGreeting(sb spanishBot) {
// 	fmt.Println(sb.getSbGreeting())
// }
// func printTbGreeting(tb thaiBot) {
// 	fmt.Println(tb.getTbGreeting())
// }
