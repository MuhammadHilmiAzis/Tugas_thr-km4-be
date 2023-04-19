package main

import () 
"fmt"

func countElements(slice []interface{}) int {
	return len(slice)
}

func main() {
	sliceInt := []int{1, 2, 3, 4, 5}
	fmt.Println(countElements(sliceInt)) // Output: 5

	sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(countElements(sliceInt)) // Output: 10

	sliceInt = []int{1, 1, 1, 5, 5, 5}
	fmt.Println(countElements(sliceInt)) // Output: 6

	sliceString := []string{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(countElements(sliceString)) // Output: 4

	sliceString = []string{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	fmt.Println(countElements(sliceString)) // Output: 8

	sliceBool := []bool{true, false, true, false, true, false}
	fmt.Println(countElements(sliceBool)) // Output: 6
}
