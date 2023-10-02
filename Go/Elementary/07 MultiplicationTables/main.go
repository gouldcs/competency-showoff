package main

import (
	"fmt"
)

func main() {
	var multiplicationTable [12][12]int = GenerateTable()
	PrintTable(multiplicationTable)
}

func GenerateTable() [12][12]int {
	var multiplicationTable [12][12]int

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			multiplicationTable[i][j] = (i + 1) * (j + 1)
		}
	}

	return multiplicationTable
}

func PrintTable(multiplicationTable [12][12]int) {
	PrintHeader()
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			currentNumber := multiplicationTable[i][j]
			if j == 0 {
				fmt.Printf("%d\t| %d", currentNumber, currentNumber)
			} else {
				fmt.Printf(", %d", currentNumber)
			}

		}
		fmt.Print("\n")
	}
}

func PrintHeader() {
	fmt.Print("*\t|")
	for i := 1; i < 13; i++ {
		if i == 1 {
			fmt.Printf(" %d", i)
		} else {
			fmt.Printf(", %d", i)
		}
	}
	fmt.Print("\n")
}
