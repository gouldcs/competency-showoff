package main

import "fmt"

/*
* | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 |
1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 |
2 | 2 | 4 | ...
3
4
5
6
7
8
9
10
11
12
*/

func main() {
	var multiplicationTable [12][12]int = generateTable()
	printTable(multiplicationTable)
}

func generateTable() [12][12]int {
	var multiplicationTable [12][12]int

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			multiplicationTable[i][j] = (i + 1) * (j + 1)
		}
	}

	return multiplicationTable
}

func printTable(multiplicationTable [12][12]int) {
	printHeader()
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			currentNumber := multiplicationTable[i][j]
			currentMax := (j + 1) * 12

			var spacing string

			if currentMax < 10 {
				spacing = "   "
			} else if currentMax < 100 {
				spacing = "  "
			} else {
				spacing = " "
			}
			if j == 0 {
				if currentNumber < 10 {
					fmt.Printf(" %d  |", currentNumber)
				} else {
					fmt.Printf(" %d |", currentNumber)
				}
			}

			fmt.Printf("%s%d%s|", spacing, currentNumber, spacing)
		}
		fmt.Print("\n")
	}
}

func printHeader() {
	fmt.Print(" *  |")
	for i := 1; i < 13; i++ {
		var currentValueMax = i * 12
		if currentValueMax >= 100 {
			fmt.Printf("   %d   |", i)
		} else if currentValueMax >= 10 {
			fmt.Printf("  %d  |", i)
		} else {
			fmt.Printf(" %d |", i)
		}
	}
	fmt.Print("\n")
}
