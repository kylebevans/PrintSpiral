package main

import (
	"fmt"
)

// print N
// for i := N-1; i > 1; i -= 2
// print i in direction
func PrintSpiral(n int) {
	fmt.Printf("\nN: %v\n", n)
	spiral := make([][]string, n)
	for i := range spiral {
		spiral[i] = make([]string, n)
	}

	direction := 0
	currentRow := 0
	currentCol := 0
	for i := 0; i < n; i++ {
		currentCol = i
		spiral[currentRow][i] = "X"
	}
	direction = 1
	for i := n - 1; i > 1; i -= 2 {
		forwardExes := currentCol + i
		backExes := currentCol - i
		upExes := currentRow - i
		downExes := currentRow + i
		for times := 0; times < 2; times++ {
			switch direction {
			case 0:
				for j := currentCol + 1; j <= forwardExes; j++ {
					currentCol = j
					spiral[currentRow][currentCol] = "X"
				}
				direction = 1
			case 1:
				for j := currentRow + 1; j <= downExes; j++ {
					currentRow = j
					spiral[currentRow][currentCol] = "X"
				}
				direction = 2
			case 2:
				for j := currentCol - 1; j >= backExes; j-- {
					currentCol = j
					spiral[currentRow][currentCol] = "X"
				}
				direction = 3
			case 3:
				for j := currentRow - 1; j >= upExes; j-- {
					currentRow = j
					spiral[currentRow][currentCol] = "X"
				}
				direction = 0
			}
		}
	}
	// handle the 1 case
	if n%2 == 0 {
		switch direction {
		case 0:
			currentCol++
			spiral[currentRow][currentCol] = "X"
		case 1:
			currentRow++
			spiral[currentRow][currentCol] = "X"
		case 2:
			currentCol--
			spiral[currentRow][currentCol] = "X"
		case 3:
			currentRow--
			spiral[currentRow][currentCol] = "X"
		}
	}
	for i := range spiral {
		for j := range spiral[i] {
			if spiral[i][j] == "" {
				fmt.Printf(" ")
			} else {
				fmt.Printf(spiral[i][j])
			}
		}
		fmt.Println()
	}
	return
}

func main() {
	for i := 1; i <= 11; i++ {
		fmt.Println()
		PrintSpiral(i)
		fmt.Println()
	}
	return
}

/*

X           1

XX          2 1
 X

XXX        3 2 2
  X
XXX
              5 4 4 2 2
XXXXX
    X
XXX X
X   X
XXXXX

XXXXXX        6 5 5 3 3 1
     X
XXXX X
X  X X
X    X
XXXXXX

XXXXXXX       7 6 6 4 4 2 2
      X
XXXXX X
X   X X
X XXX X
X     X
XXXXXXX
              8 7 7 5 5 3 3 1
XXXXXXXX
       X
XXXXXX X
X    X X
X X X X
X XXXX X
X      X
XXXXXXXX
              9 8 8 6 6 4 4 2 2
XXXXXXXXX
        X
XXXXXXX X
X     X X
X XXX X X
X X   X X
X XXXXX X
X       X
XXXXXXXXX

*/
