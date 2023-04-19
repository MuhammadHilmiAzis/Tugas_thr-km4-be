package main

import "fmt"

func addElement(array []int, element int, position string) []int {
	if position == "up" {
		return append([]int{element}, array...)
	} else if position == "down" {
		return append(array, element)
	} else {
		return array
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	element := 6

	newArray := addElement(array, element, "up")
	fmt.Println(newArray) // Output: [6 1 2 3 4 5]

	newArray = addElement(array, element, "down")
	fmt.Println(newArray) // Output: [1 2 3 4 5 6]

	newArray = addElement(array, element, "invalid")
	fmt.Println(newArray) // Output: [1 2 3 4 5]
}
