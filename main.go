package main

import (
	"first-go/contracts"
	"first-go/cryptography"
	"first-go/goroutine"
	"first-go/operations"
	"first-go/services"
	"first-go/structs"
	"fmt"
	"time"
)

func main() {

	// todo: palindrome check
	checkString := []string{"level", "racecar", "hello", "world", "1232"}
	fmt.Println("Palindrome check:")
	for _, str := range checkString {
		fmt.Printf("%s: %t\n", str, operations.IsPalindrome(str))
	}

	// todo: fibonacci series
	checkNumber := []int{0, 1, 3, 2, 5, 4, 1}
	fmt.Println("Fibonacci series:")
	for _, num := range checkNumber {
		fmt.Printf("%d: %d\n", num, operations.Fibonacci(num))
	}

	// todo: prime number check
	checkPrime := []int{2, 3, 4, 5, 7, 10, 11, 20, 23, 24}
	fmt.Println("Prime number check:")
	for _, num := range checkPrime {
		fmt.Printf("%d: %t\n", num, operations.IsPrime(num))
	}

	// todo: odd/even number check
	checkNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Odd/even number check:")
	for _, num := range checkNumbers {
		fmt.Printf("%d: Even: %t, Odd: %t\n", num, operations.IsEven(num), operations.IsOdd(num))
	}

	// todo: print stars
	fmt.Println("Print stars:")
	operations.PrintStars(5)

	// todo: factorial calculation
	checkFactorial := []int{0, 1, 3, 4, 5, 6}
	fmt.Println("Factorial calculation:")
	for _, num := range checkFactorial {
		fmt.Printf("%d: %d\n", num, operations.Factorial(num))
	}

	// todo: string manipulation
	str := "Hello, World!"
	fmt.Println("String manipulation:")
	fmt.Printf("Original string: %s\n", str)
	fmt.Printf("Reversed string: %s\n", operations.ReverseString(str))
	fmt.Printf("Uppercase string: %s\n", operations.ToUpperCase(str))
	fmt.Printf("Lowercase string: %s\n", operations.ToLowerCase(str))
	fmt.Printf("Replace spaces: %s\n", operations.ReplaceSpaces(str, "_"))

	// todo: array operations
	fmt.Println("Array operations:")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	operations.ArrayOperations(arr)

	// todo: pointer
	fmt.Println("Pointer:")
	a, b := 10, 20
	fmt.Printf("Original values: a = %d, b = %d\n", a, b)

	operations.Swap(&a, &b)
	fmt.Printf("Swapped values: a = %d, b = %d\n", a, b)

	operations.Increment(&a)
	fmt.Printf("Incremented value: a = %d\n", a)

	operations.Decrement(&b)
	fmt.Printf("Decremented value: b = %d\n", b)

	operations.PrintPointerValue(&a)
	operations.PrintPointerValue(&b)

	// todo: struct
	fmt.Println("Struct:")
	person := structs.NewPerson("Alice", 30)
	person.Great()

	person.HaveBirthday()
	person.Great()

	// todo: interface
	fmt.Println("Interface:")
	shapes := []contracts.Shape{
		services.Rectangle{Width: 5, Height: 10},
		services.Circle{Radius: 3},
	}

	for _, shape := range shapes {
		fmt.Printf("Shape: %T\n", shape)
		fmt.Printf("Area: %f\n", shape.Area())
		fmt.Printf("Perimeter: %f\n", shape.Perimeter())
	}

	// todo: goroutine
	fmt.Println("Goroutine:")
	go goroutine.PrintNumbers()
	go goroutine.PrintLetters()

	time.Sleep(time.Second * 6)
	fmt.Println("Main function finished.")

	// todo: encryption/decryption
	text := "ATTACTATONCE"
	shift := 4
	encryptedText := cryptography.Encrypt(text, int8(shift))
	decryptedText := cryptography.Decrypt(encryptedText, int8(shift))

	fmt.Printf("Original text: %s\n", text)
	fmt.Printf("Shift: %d\n", shift)
	fmt.Printf("Encrypted text: %s\n", encryptedText)
	fmt.Printf("Decrypted text: %s\n", decryptedText)
}
