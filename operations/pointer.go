package operations

import "fmt"

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func Increment(a *int) {
	*a++
}

func Decrement(a *int) {
	*a--
}

func PrintPointerValue(a *int) {
	fmt.Println("Pointer value:", *a)
}
