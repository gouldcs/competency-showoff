package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var guesses hashset.Set = *hashset.New()
	var randomNumber int = r1.Intn(100)
	var isCorrect bool = false
	var lastGuess int

	for !isCorrect {
		fmt.Printf("Guess a number (Guess number %d): ", guesses.Size()+1)
		fmt.Scanf("%d\n", &lastGuess)

		if !guesses.Contains(lastGuess) {
			guesses.Add(lastGuess)
		}

		if lastGuess == randomNumber {
			fmt.Println("Correct!")
			fmt.Printf("It only took you %d guesses!", guesses.Size())
			isCorrect = true
		} else if randomNumber > lastGuess {
			fmt.Println("The number you're trying to guess is higher.")
		} else if randomNumber < lastGuess {
			fmt.Println("The number you're trying to guess is lower.")
		}
	}
}
