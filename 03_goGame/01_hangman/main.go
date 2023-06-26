package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	"unicode"
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
	targetWord := getRandomWord()
	guessedLetters := initializeGuessdWords(targetWord)
	// targetWord = "United States of America"
	fmt.Println(targetWord)
	guessedLetters['s'] = true
	hangmanState := 1
	printGameState(targetWord, guessedLetters, hangmanState)
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

func initializeGuessdWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true
	return guessedLetters
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

func printGameState(
	targetWrod string,
	guessedLetters map[rune]bool,
	hangmanState int,
) {

	fmt.Println(getWordGuessingProgress(targetWrod, guessedLetters))
	fmt.Println()
	fmt.Println(printHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(
	targetWord string,
	guessedLetters map[rune]bool,
) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}
		result += " "
	}
	return result

}

func printHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}

	return string(data)
}
