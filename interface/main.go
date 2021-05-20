package main

import "fmt"

// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct {
// }
// type spanishBot struct {
// }
// type thaiBot struct {
// }

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		base:   10,
		height: 20,
	}
	printArea(t)
	s := square{
		sideLength: 10,
	}
	printArea(s)
	// resp, err := http.Get("https://www.google.com")
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)

	// fmt.Println(resp.Status)
	// fmt.Println(resp.Header["Content-Type"])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(body)
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

func (t triangle) getArea() float64 {
	return (1 / 2) * t.base * t.height
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func printArea(sh shape) {
	fmt.Println(sh.getArea())
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
