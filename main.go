package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)

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
	// targetWord := "United States of America"
	guessedLetters := initializeGuessWord(targetWord)
	hangmanState := 0

	for {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. Please use letters only...")
			continue
		}
	}

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

func initializeGuessWord(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]

	return targetWord
}

func printGameState(
	targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int,
) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(
	targetWord string,
	guessedLetters map[rune]bool,
) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}
		result += " "
	}

	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(
		fmt.Sprintf("states/hangman%d", hangmanState),
	)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	fmt.Print("> ")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}
