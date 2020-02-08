package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.saveToFile("test.txt")
	cards := newDeckFromFile("test.txt")
	fmt.Println(cards.toString())
	cards.shuffle()
	fmt.Println("--------------")
	fmt.Println(cards.toString())
}
