package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"United States of America",
	"Indonesia",
	"Nazism",
	"Apple",
	"Programming",
}

func main() {
	// Derive a word we have to guess
	rand.Seed(time.Now().UnixNano())
	targetWord := dictionary[rand.Intn(len(dictionary))]
	fmt.Println(targetWord)

	// Printing game state
	//  * Print word you're guessing
	//  * Print hangman state

	// Derive a word we have to guess
	// Read user input
	//   * validate it (eg. only letters)
	// Determine if the letter is a correct guess or not
	//   * if correct, update the guesed letters
	//   * if incorrect, update the hangman state
	// if word is guessed -> game over, you win
	// If hangman is complete  -> game over, you lose
	// then do that
}
