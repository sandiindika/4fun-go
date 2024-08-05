package operations

import "fmt"

func ArrayOperations(arr []int) {

	duplicateArr := []int{1, 2, 3, 4, 5}
	duplicateAgain := [...]int{6, 7, 8, 9, 10}

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("duplicateArr: %v\n", duplicateArr)
	fmt.Printf("duplicateAgain: %v\n", duplicateAgain)

	arr[0] = 10
	arr[1] = 20
	fmt.Printf("Modified arr: %v\n", arr)

	fmt.Println("Iterating over duplicateArr")
	for i, v := range duplicateArr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	var duplicateCopy [5]int
	copy(duplicateCopy[:], arr[:])
	fmt.Printf("Copied duplicateCopy from arr: %v\n", duplicateCopy)

	fmt.Printf("Length of duplicateAgain: %d\n", len(duplicateAgain))
}
