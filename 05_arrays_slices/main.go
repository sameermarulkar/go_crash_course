package main

import (
	"fmt"
)

func main() {
	// Arrays
	// var fruitArr [2]string - arrays have to be fixed lengths

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Decalre and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}	// slices == variable lengthed arrays. note the curly braces

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
