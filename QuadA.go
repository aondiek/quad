package main

import "fmt"

func QuadA(x, y int) {
	/*if either width or height is 0 or negative, do nothing*/
	if x <= 0 || y <= 0 {
		return
	}
	/*loop over each row, from top to bottom*/
	for i := 0; i < y; i++ {
		/*loop over each column in the current row, from left to right*/
		for j := 0; j < x; j++ {
			switch {
			/*top corners, if it is the first row and first or last column, print '0'*/
			case (i == 0 && (j == 0 || j == x-1)):
				fmt.Print("0")
			/*bottom corners, if it's the last row and first or last column, print '0'*/
			case (i == y-1 && (j == 0 || j == x-1)):
				fmt.Print("0")
			/*edges excluding corners, print '-'*/
			case (i == 0 || i == y-1 || j == 0 || j == x-1):
				fmt.Print("-")
			/*inside of the rectangle, print a space*/
			default:
				fmt.Print(" ")
			}
		}
		/*move to the next line*/
		fmt.Println()
	}
}
