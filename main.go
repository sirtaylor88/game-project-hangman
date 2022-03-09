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

const MaxPossibleHints = 1

func main() {
	rand.Seed(time.Now().UnixNano())

	targetWord := getRandomWord()
	guessedLetters := initializeGuessWord(targetWord)
	hangmanState := 0
	usedLetters := []rune{}
	hintCount := 0

	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState, hintCount)
		input := readInput()
		if len(input) != 1 {
			if input == "hint" && hintCount <= MaxPossibleHints {
				fmt.Printf(
					"Hint: There is a letter %c in this word!\n\n",
					unicode.ToLower(hint(targetWord, guessedLetters)),
				)
				hintCount++
			} else if input == "hint" {
				fmt.Println("You already used all available hints!")
			} else {
				fmt.Println("Invalid input. Please use letters only...")

			}
			continue

		}

		letter := rune(input[0])

		if contains(usedLetters, unicode.ToLower(letter)) {
			fmt.Printf("You've already used the letter %c\n", letter)
		} else {
			usedLetters = append(usedLetters, unicode.ToLower(letter))
		}

		if isCorrectGuess(targetWord, letter) {
			guessedLetters[unicode.ToLower(letter)] = true
		} else {
			hangmanState++
		}
	}

	printGameState(targetWord, guessedLetters, hangmanState, hintCount)
	fmt.Print("Game Over...")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You lose!")
	} else {
		panic("Invalid state. Game is over and there is no winner!")
	}
}

func contains(s []rune, r rune) bool {
	for _, el := range s {
		if el == r {
			return true
		}
	}
	return false
}

func initializeGuessWord(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}

	index1, index2 := rand.Intn(len(targetWord)-1), rand.Intn(len(targetWord)-1)
	guessedLetters[unicode.ToLower(rune(targetWord[index1]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[index2]))] = true

	return guessedLetters
}

func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]

	return targetWord
}

func isGameOver(
	targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int,
) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}

	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func printGameState(
	targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int,
	hintCount int,
) {
	fmt.Println("\n***********************")
	fmt.Println()
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	if hintCount < MaxPossibleHints {
		fmt.Println("If you want a hint, enter 'hint'")
	}
	fmt.Printf("Hints used: %d/%d \n", hintCount, MaxPossibleHints)
	fmt.Println(getHangmanDrawing(hangmanState))
	fmt.Println("\n***********************")
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

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(strings.ToLower(targetWord), unicode.ToLower(letter))
}

func hint(targetWord string, guessedLetters map[rune]bool) rune {
	unguessedLetters := []rune{}
	for _, ch := range targetWord {
		if _, ok := guessedLetters[unicode.ToLower(rune(ch))]; !ok {
			unguessedLetters = append(unguessedLetters, rune(ch))
		}
	}

	return unguessedLetters[rand.Intn(len(unguessedLetters)-1)]
}
