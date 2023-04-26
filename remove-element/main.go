package main

import "fmt"

func removeElement(arr []interface{}, position string) []interface{} {
	if position == "left" {
		return arr[1:]
	} else if position == "right" {
		return arr[:len(arr)-1]
	} else {
		panic("Invalid position value, position should be either 'left' or 'right'")
	}
}

func main() {
	arr1 := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(removeElement(arr1, "left")) // output: [2 3 4 5]

	arr2 := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(removeElement(arr2, "right")) // output: [1 2 3 4]

	arr3 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeElement(arr3, "left")) // output: [Budi Joko Tono]

	arr4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeElement(arr4, "right")) // output: [Edo Budi Joko]
}
