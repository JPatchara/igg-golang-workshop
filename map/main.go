package main

import "fmt"

func main() {
	ages := map[string]int16{
		"Jin":  25,
		"Jone": 35,
		"Juy":  99,
	}
	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#00FF00",
	// 	"blue":  "#0000FF",
	// }

	// colors := make(map[string]string)
	// colors["red"] = "#FF0000"
	// colors["green"] = "#00FF00"
	// colors["blue"] = "#0000FF"

	fmt.Println(ages)
	// updateMap(colors, "black", "#000000")
	// printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hexCode := range colors {
		fmt.Println(color, hexCode)
	}
}

func updateMap(colors map[string]string, color string, code string) {
	colors[color] = code
}
