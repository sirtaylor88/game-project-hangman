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
	"Gazelle",
	"Apple",
	"Programming",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	targetWord := getRandomWord()
	fmt.Println(targetWord)
	// Printing game state
	//   *  Print word you're guessing
	//   *  Print hangman state
	// Read user input
	//   *  Validate it (e.g. only letters)
	// Determine if the letter is a correct guess or not
	//   *  If correct, update the guessed letter
	//   *  If incorrect, update the hangman state
	// If word is guessed, game over, you win
	// If hangman is complete, game over, you lose
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}
