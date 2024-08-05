package goroutine

import (
	"fmt"
	"time"
)

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func PrintLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("Letter: %c\n", i)
		time.Sleep(1 * time.Second)
	}
}
